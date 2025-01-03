package config

import "time"

var DefaultConfig = map[string]any{
	"app": map[string]any{
		"debug":       false,
		"version":     "0.0.1",
		"environment": "production",
	},
	"http": map[string]any{
		"port":         8080,
		"idleTimeout":  "60s",
		"readTimeout":  "5s",
		"writeTimeout": "10s",
	},
	"database": map[string]any{
		"dsn":                "postgres://username:password@host:port/database",
		"maxOpenConnections": 25,
		"maxIdleConnections": 25,
		"maxIdleTime":        "15m",
	},
	"authentication": map[string]any{
		"accessSecretKey":       "access_secret_key",
		"refreshSecretKey":      "refresh_secret_key",
		"accessTokenExpiresIn":  "5m",
		"refreshTokenExpiresIn": "72h",
	},
}

type Configuration struct {
	App            AppConfig            `mapstructure:"app"`
	Http           HttpServerConfig     `mapstructure:"http"`
	Database       DatabaseConfig       `mapstructure:"database"`
	Authentication AuthenticationCofnig `mapstructure:"authentication"`
}

type AppConfig struct {
	Debug       bool   `mapstructure:"debug"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

type HttpServerConfig struct {
	Port         int           `mapstructure:"port"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

type DatabaseConfig struct {
	Dsn                string        `mapstructure:"dsn"`
	MaxOpenConnections int           `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int           `mapstructure:"maxIdleConnections"`
	MaxIdleTime        time.Duration `mapstructure:"maxIdleTime"`
}

type AuthenticationCofnig struct {
	AccessSecretKey       string        `mapstructure:"accessSecretKey"`
	RefreshSecretKey      string        `mapstructure:"refreshSecretKey"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"accessTokenExpiresIn"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"refreshTokenExpiresIn"`
}
