package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/server/handlers/health"
	"github.com/SIProjects/faucet-api/server/handlers/payouts"
	"github.com/SIProjects/faucet-api/server/handlers/queue"
	"github.com/SIProjects/faucet-api/server/system"
	"github.com/gorilla/mux"
)

type Server struct {
	System *system.System
	Router *mux.Router
}

func New(db database.Database, c cache.Cache, ch *chain.Chain) (*Server, error) {
	r := mux.NewRouter()
	s := Server{
		System: system.New(db, c, ch, r),
		Router: r,
	}
	s.SetupRoutes()
	return &s, nil
}

func (s *Server) SetupRoutes() {
	health.SetupRoutes(s.System)
	queue.SetupRoutes(s.System)
	payouts.SetupRoutes(s.System)
}

func (s *Server) Start(done chan struct{}) {
	wait := time.Second * 15

	srv := &http.Server{
		Addr:         ":3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.System.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Server listening on port 3000")

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	go func() {
		srv.Shutdown(ctx)
	}()
	<-ctx.Done()
}
