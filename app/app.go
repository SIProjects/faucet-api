package app

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/scheduler"
	"github.com/SIProjects/faucet-api/server"
)

type App struct {
	Logger          *log.Logger
	DB              *database.PGDatabase
	Cache           cache.Cache
	Chain           *chain.Chain
	Server          *server.Server
	PayoutScheduler scheduler.Scheduler
	done            chan struct{}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(
	db *database.PGDatabase,
	c cache.Cache,
	ch *chain.Chain,
	sch scheduler.Scheduler,
	l *log.Logger,
) (*App, error) {
	s, err := server.New(db, c, ch, l)
	if err != nil {
		return nil, err
	}
	a := App{
		DB:              db,
		Cache:           c,
		Chain:           ch,
		Server:          s,
		PayoutScheduler: sch,
		Logger:          l,
	}
	return &a, nil
}

func (a *App) Start() {
	a.done = make(chan struct{})
	done := make(chan struct{})

	go a.Server.Start(done)
	go a.PayoutScheduler.Start(done)

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

	a.Logger.Println("shutting down")
}

func (a *App) Close() {
	a.Logger.Println("Closing app")
	a.DB.Close()
	a.Cache.Close()
	a.Logger.Println("App closed")
}
