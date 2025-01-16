package internal

import (
	"context"
	"os"
	"path/filepath"

	"github.com/asztemborski/syncro/internal/api"
	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/asztemborski/syncro/internal/store"
	"go.uber.org/zap"
)

type BootstrapArgs struct {
	ConfigFile string
	EnvFile    string
	EnvPrefix  string
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
		config.WithEnvPrefix(args.EnvPrefix),
	)

	cfg, err := loader.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	db, err := store.NewPostgresDb(cfg.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}

	store := store.NewSqlStore(db)
	accountService := app.NewAccountService(store.Account())

	app := app.NewApp(cfg, logger, accountService)
	server := api.NewServer(app)
	server.Start(ctx)
}
