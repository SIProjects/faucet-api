package queue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SIProjects/faucet-api/server/system"
	"github.com/btcsuite/btcutil"
)

type AddBody struct {
	Address string `json:"address"`
}

func AddToQueue(s *system.System) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var body AddBody

		err = json.Unmarshal(data, &body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		addr, err := s.Chain.DecodeAddress(body.Address)

		if err != nil {
			http.Error(w, "InvalidAddress", http.StatusBadRequest)
			return
		}

		can, err := s.Cache.CanAddAddress(addr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !can {
			http.Error(w, "AddressAlreadyQueued", http.StatusBadRequest)
			return
		}

		queueLen, err := s.Cache.GetQueuedCount()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		balance, err := s.Chain.Node.GetBalance()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		requiredBalance := (queueLen + int64(1)) * int64(s.Config.MaxPayout)

		required, err := btcutil.NewAmount(float64(requiredBalance))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if balance < required {
			http.Error(
				w,
				"Not enough balance in the faucet",
				http.StatusInternalServerError,
			)
			return
		}

		err = s.Cache.AddAddressToQueue(addr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
	}
}
