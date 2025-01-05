package app

import (
	"github.com/asztemborski/syncro/internal/config"
	"go.uber.org/zap"
)

type App struct {
	config *config.Configuration
	logger *zap.Logger
}

func NewApp(config *config.Configuration, logger *zap.Logger) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (a *App) Config() *config.Configuration {
	return a.config
}

func (a *App) Logger() *zap.Logger {
	return a.logger
}
