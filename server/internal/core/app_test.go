package core_test

import (
	"testing"

	"github.com/asztemborski/syncro/internal/core"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestNewApp(t *testing.T) {
	app := core.NewApp(&config.Configuration{}, zaptest.NewLogger(t), &core.AccountService{})
	assert.NotNil(t, app.Config())
}
