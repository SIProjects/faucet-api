package queue

import (
	"github.com/SIProjects/faucet-api/constants"
	"github.com/SIProjects/faucet-api/server/system"
)

func SetupRoutes(s *system.System) {
	s.Router.HandleFunc(
		"/v1/queue", AddToQueue(s),
	).Methods("POST").HeadersRegexp(
		"Content-Type", constants.APPLICATION_JSON_REGEX,
	)
}
