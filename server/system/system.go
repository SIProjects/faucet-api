package system

import (
	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/gorilla/mux"
)

type System struct {
	DB     database.Database
	Cache  cache.Cache
	Chain  *chain.Chain
	Router *mux.Router
}

func New(db database.Database, c cache.Cache, ch *chain.Chain, r *mux.Router) *System {
	return &System{
		DB:     db,
		Cache:  c,
		Chain:  ch,
		Router: r,
	}
}
