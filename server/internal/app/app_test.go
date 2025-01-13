package app_test

import (
	"testing"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestNewApp(t *testing.T) {
	app := app.NewApp(&config.Configuration{}, zaptest.NewLogger(t), &app.AccountService{})
	assert.NotNil(t, app.Config())
}
