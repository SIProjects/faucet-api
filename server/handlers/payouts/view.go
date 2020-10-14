package payouts

import (
	"encoding/json"
	"net/http"

	"github.com/SIProjects/faucet-api/resources/payout"
	"github.com/SIProjects/faucet-api/server/system"
)

func ReadPayouts(s *system.System) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payouts, err := payout.Recent(s.DB)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(payouts)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
