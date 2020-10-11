package health

import "github.com/SIProjects/faucet-api/server/system"

func SetupRoutes(s *system.System) {
	s.Router.Methods("GET").Path("/").HandlerFunc(GetHealth)
}
