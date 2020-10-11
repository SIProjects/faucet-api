package app

import "github.com/SIProjects/faucet-api/database"

type App struct {
}

func New(db *database.Database) (*App, error) {
	a := App{}
	return &a, nil
}

func (a *App) Start() {
}
