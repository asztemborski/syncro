package app

import (
	"github.com/asztemborski/syncro/internal/config"
	"go.uber.org/zap"
)

type App struct {
	config         *config.Configuration
	logger         *zap.Logger
	accountService *AccountService
}

func NewApp(config *config.Configuration, logger *zap.Logger, accountService *AccountService) *App {
	return &App{
		config:         config,
		logger:         logger,
		accountService: accountService,
	}
}

func (a *App) Config() *config.Configuration {
	return a.config
}

func (a *App) Logger() *zap.Logger {
	return a.logger
}

func (a *App) AccountService() *AccountService {
	return a.accountService
}
