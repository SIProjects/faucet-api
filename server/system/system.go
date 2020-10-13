package system

import (
	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/node"
	"github.com/gorilla/mux"
)

type System struct {
	DB     database.Database
	Cache  cache.Cache
	Node   node.Node
	Router *mux.Router
}

func New(db database.Database, c cache.Cache, n node.Node, r *mux.Router) *System {
	return &System{
		DB:     db,
		Cache:  c,
		Node:   n,
		Router: r,
	}
}
