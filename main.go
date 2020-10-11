package main

import (
	"log"

	"github.com/SIProjects/faucet-api/app"
)

func main() {
	a, err := app.New()

	if err != nil {
		log.Fatalln("Error starting app:", err)
	}

	a.Start()
}
