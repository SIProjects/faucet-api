package queue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SIProjects/faucet-api/server/system"
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

		err = s.Cache.AddAddressToQueue(addr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
	}
}
