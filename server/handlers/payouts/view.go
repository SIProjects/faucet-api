package payouts

import (
	"net/http"

	"github.com/SIProjects/faucet-api/server/system"
)

func ReadPayouts(s *system.System) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
