package app

import "github.com/SIProjects/faucet-api/database"

type App struct {
	DB database.Database
}

func New(db database.Database) (*App, error) {
	a := App{
		DB: db,
	}
	return &a, nil
}

func (a *App) Start() {
}
