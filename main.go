package main

import (
	"log"
	"os"
	"time"

	"github.com/SIProjects/faucet-api/app"
	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/node"
	"github.com/SIProjects/faucet-api/scheduler"

	_ "github.com/joho/godotenv/autoload"
)

func loadApp() (*app.App, error) {
	db, err := database.New(os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	node, err := node.New(
		os.Getenv("RPC_URL"),
		os.Getenv("RPC_USER"),
		os.Getenv("RPC_PASSWORD"),
	)

	if err != nil {
		return nil, err
	}

	ch := chain.New(node, chain.Testnet)

	c, err := cache.New(
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_NAME"),
		ch,
	)

	if err != nil {
		return nil, err
	}

	sch, err := scheduler.New(time.Second*5, db, c, ch)

	a, err := app.New(db, c, ch, sch)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func main() {
	a, err := loadApp()
	if err != nil {
		log.Fatalln("Error loading app:", err)
	}
	a.Start()
}
