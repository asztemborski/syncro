package app

import "github.com/asztemborski/syncro/internal/config"

type App struct {
	config *config.Configuration
}

func NewApp(config *config.Configuration) *App {
	return &App{
		config: config,
	}
}

func (a *App) Config() *config.Configuration {
	return a.config
}
