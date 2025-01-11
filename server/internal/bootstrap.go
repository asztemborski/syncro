package internal

import (
	"context"
	"os"
	"path/filepath"

	"github.com/asztemborski/syncro/internal/api"
	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"go.uber.org/zap"
)

type BootstrapArgs struct {
	ConfigFile string
	EnvFile    string
}

func Run(ctx context.Context, args BootstrapArgs) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ex, err := os.Executable()
	if err != nil {
		logger.Fatal(err.Error())
	}

	appDir := filepath.Dir(ex)
	loader := config.NewLoader(
		config.WithConfigFile(filepath.Join(appDir, args.ConfigFile)),
		config.WithEnvFile(filepath.Join(appDir, args.EnvFile)),
		config.WithDefaults(config.DefaultConfig),
	)

	cfg, err := loader.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	app := app.NewApp(cfg, logger)
	server := api.NewServer(app)
	server.Start(ctx)
}
