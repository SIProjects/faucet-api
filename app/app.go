package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/node"
	"github.com/SIProjects/faucet-api/server"
)

type App struct {
	DB     database.Database
	Node   *node.Node
	Server *server.Server
}

func New(db database.Database, n *node.Node) (*App, error) {
	s, err := server.New(db, n)
	if err != nil {
		return nil, err
	}
	a := App{
		DB:     db,
		Node:   n,
		Server: s,
	}
	return &a, nil
}

func (a *App) Start() {
	done := make(chan struct{})

	go a.Server.Start(done)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive interrupt signal.
	<-c

	done <- struct{}{}

	log.Println("shutting down")
}
