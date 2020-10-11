package system

import (
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/node"
	"github.com/gorilla/mux"
)

type System struct {
	DB     database.Database
	Node   *node.Node
	Router *mux.Router
}

func New(db database.Database, n *node.Node, r *mux.Router) *System {
	return &System{
		DB:     db,
		Node:   n,
		Router: r,
	}
}
