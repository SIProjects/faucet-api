package app

type App struct {
}

func New() (*App, error) {
	a := App{}
	return &a, nil
}

func (a *App) Start() {
}
