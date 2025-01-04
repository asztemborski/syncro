package internal

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/asztemborski/syncro/internal/api"
	"github.com/asztemborski/syncro/internal/api/handler"
	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
)

type BootstrapArgs struct {
	ConfigFile string
	EnvFile    string
}

func Run(ctx context.Context, args BootstrapArgs) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	appDir := filepath.Dir(ex)
	loader := config.NewLoader(
		config.WithConfigFile(filepath.Join(appDir, args.ConfigFile)),
		config.WithEnvFile(filepath.Join(appDir, args.EnvFile)),
		config.WithDefaults(config.DefaultConfig),
	)

	cfg, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(cfg)
	server := api.NewServer(app)
	server.RegisterHandlers(
		handler.NewHealthHandler(app),
	)

	server.Start(ctx)
}
