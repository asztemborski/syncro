package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/asztemborski/syncro/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	loader := config.NewLoader(config.WithDefaults(config.DefaultConfig))
	cfg, err := loader.Load()
	assert.NoError(t, err, "should not return an error")

	assert.Equal(t, false, cfg.App.Debug)
	assert.Equal(t, "0.0.1", cfg.App.Version)
	assert.Equal(t, "production", cfg.App.Environment)

	assert.Equal(t, 8080, cfg.Http.Port)
	assert.Equal(t, 60*time.Second, cfg.Http.IdleTimeout)
	assert.Equal(t, 5*time.Second, cfg.Http.ReadTimeout)
	assert.Equal(t, 10*time.Second, cfg.Http.WriteTimeout)

	assert.Equal(t, "postgres://username:password@host:port/database", cfg.Database.Dsn)
	assert.Equal(t, 25, cfg.Database.MaxOpenConnections)
	assert.Equal(t, 25, cfg.Database.MaxIdleConnections)
	assert.Equal(t, 15*time.Minute, cfg.Database.MaxIdleTime)

	assert.Equal(t, "access_secret_key", cfg.Authentication.AccessSecretKey)
	assert.Equal(t, "refresh_secret_key", cfg.Authentication.RefreshSecretKey)
	assert.Equal(t, 5*time.Minute, cfg.Authentication.AccessTokenExpiresIn)
	assert.Equal(t, 72*time.Hour, cfg.Authentication.RefreshTokenExpiresIn)
}

func TestLoadConfigFile(t *testing.T) {
	fileContent := `
app:
  debug: true
  version: "1.0.0"
http:
  port: 9090
`
	os.WriteFile("test_config.yaml", []byte(fileContent), 0644)
	defer os.Remove("test_config.yaml")

	loader := config.NewLoader(
		config.WithConfigFile("test_config.yaml"),
		config.WithDefaults(config.DefaultConfig),
	)

	cfg, err := loader.Load()
	assert.NoError(t, err)

	assert.Equal(t, true, cfg.App.Debug)
	assert.Equal(t, "1.0.0", cfg.App.Version)
	assert.Equal(t, 9090, cfg.Http.Port)
	assert.Equal(t, "production", cfg.App.Environment)
}

func TestLoadEnvFile(t *testing.T) {
	envFileContent := `authentication.accessSecretKey=super_secret_key`
	os.WriteFile(".env", []byte(envFileContent), 0644)
	defer os.Remove(".env")

	loader := config.NewLoader(
		config.WithEnvFile(".env"),
		config.WithDefaults(config.DefaultConfig),
		config.WithEnvPrefix("APP"),
	)

	cfg, err := loader.Load()
	assert.NoError(t, err)

	assert.Equal(t, cfg.Authentication.AccessSecretKey, "super_secret_key")
}

func TestLoadWithDefaults(t *testing.T) {
	loader := config.NewLoader(
		config.WithDefaults(map[string]any{
			"app": map[string]any{
				"debug": true,
			},
		}),
	)

	cfg, err := loader.Load()
	assert.NoError(t, err)
	assert.Equal(t, true, cfg.App.Debug)
}

func TestInvalidConfigFile(t *testing.T) {
	loader := config.NewLoader(config.WithConfigFile("invalid_file.yaml"))

	_, err := loader.Load()
	assert.Error(t, err)
}

func TestUnmarshalError(t *testing.T) {
	os.WriteFile("bad_config.yaml", []byte(`app: "invalid_structure"`), 0644)
	defer os.Remove("bad_config.yaml")

	loader := config.NewLoader(config.WithConfigFile("bad_config.yaml"))
	_, err := loader.Load()
	assert.Error(t, err)
}
