package app_test

import (
	"testing"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := app.NewApp(&config.Configuration{})
	assert.NotNil(t, app.Config())
}
