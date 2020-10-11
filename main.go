package main

import (
	"log"
	"os"

	"github.com/SIProjects/faucet-api/app"
	"github.com/SIProjects/faucet-api/database"

	_ "github.com/joho/godotenv/autoload"
)

func loadApp() (*app.App, error) {
	db, err := database.New(os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	a, err := app.New(db)

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
