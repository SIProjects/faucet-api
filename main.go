package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SIProjects/faucet-api/app"
	"github.com/SIProjects/faucet-api/cache"
	"github.com/SIProjects/faucet-api/chain"
	"github.com/SIProjects/faucet-api/configuration"
	"github.com/SIProjects/faucet-api/database"
	"github.com/SIProjects/faucet-api/node"
	"github.com/SIProjects/faucet-api/scheduler"

	_ "github.com/joho/godotenv/autoload"
)

func loadIntVar(name string, defaultValue int) int {
	val, err := strconv.Atoi(os.Getenv(name))

	if err != nil {
		return defaultValue
	}

	return val
}

func loadApp() (*app.App, error) {
	logger := log.New(
		os.Stdout,
		"faucet-api: ",
		log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Lshortfile,
	)

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

	minPayout := loadIntVar("MIN_PAYOUT", 10)
	maxPayout := loadIntVar("MAX_PAYOUT", 100)

	config := configuration.New(minPayout, maxPayout)

	sch, err := scheduler.New(
		time.Second*30, db, c, ch, logger, config,
	)

	a, err := app.New(db, c, ch, sch, logger, config)

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
