package payouts

import (
	"github.com/SIProjects/faucet-api/constants"
	"github.com/SIProjects/faucet-api/server/system"
)

func SetupRoutes(s *system.System) {
	s.Router.HandleFunc(
		"/v1/payouts", CreatePayout(s),
	).Methods("POST").HeadersRegexp(
		"Content-Type", constants.APPLICATION_JSON_REGEX,
	)

	s.Router.HandleFunc(
		"/v1/payouts", ReadPayouts(s),
	).Methods("GET")
}
