package payouts

import (
	"github.com/SIProjects/faucet-api/server/system"
)

func SetupRoutes(s *system.System) {
	s.Router.HandleFunc(
		"/v1/payouts", ReadPayouts(s),
	).Methods("GET")
}
