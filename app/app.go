package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/server"
)

type App struct {
	DB     *database.PGDatabase
	Cache  cache.Cache
	Chain  *chain.Chain
	Server *server.Server
	done   chan struct{}
}

func New(db *database.PGDatabase, c cache.Cache, ch *chain.Chain) (*App, error) {
	s, err := server.New(db, c, ch)
	if err != nil {
		return nil, err
	}
	a := App{
		DB:     db,
		Cache:  c,
		Chain:  ch,
		Server: s,
	}
	return &a, nil
}

func (a *App) Start() {
	a.done = make(chan struct{})
	done := make(chan struct{})

	go a.Server.Start(done)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive interrupt signal.
	select {
	case <-c:
		break
	case <-a.done:
		break
	}

	done <- struct{}{}

	log.Println("shutting down")
}

func (a *App) Close() {
	log.Println("Closing app")
	a.DB.Close()
	a.Cache.Close()
	log.Println("App closed")
}
