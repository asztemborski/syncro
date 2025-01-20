package config

import "time"

var DefaultConfig = map[string]any{
	"core": map[string]any{
		"debug":       false,
		"version":     "0.0.1",
		"environment": "production",
	},
	"http": map[string]any{
		"port":         8080,
		"idleTimeout":  "60s",
		"readTimeout":  "5s",
		"writeTimeout": "10s",
		"rateLimiter": map[string]any{
			"enabled":      true,
			"maximumRPS":   20,
			"maximumBurst": 10,
		},
	},
	"database": map[string]any{
		"dsn":                "postgres://username:password@host:port/database",
		"maxOpenConnections": 25,
		"maxIdleConnections": 25,
		"maxIdleTime":        "15m",
		"maxOpenConns":       25,
		"maxIdleConns":       25,
	},
	"authentication": map[string]any{
		"accessSecretKey":       "access_secret_key",
		"refreshSecretKey":      "refresh_secret_key",
		"accessTokenExpiresIn":  "5m",
		"refreshTokenExpiresIn": "72h",
	},
}

type Configuration struct {
	App            AppConfig            `mapstructure:"core"`
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
	Port         int               `mapstructure:"port"`
	IdleTimeout  time.Duration     `mapstructure:"idleTimeout"`
	ReadTimeout  time.Duration     `mapstructure:"readTimeout"`
	WriteTimeout time.Duration     `mapstructure:"writeTimeout"`
	RateLimiter  RateLimiterConfig `mapstructure:"rateLimiter"`
}

type RateLimiterConfig struct {
	MaximumRPS   float64 `mapstructure:"maximumRPS"`
	MaximumBurst int     `mapstructure:"maximumBurst"`
	Enabled      bool    `mapstructure:"enabled"`
}

type DatabaseConfig struct {
	Dsn                string        `mapstructure:"dsn"`
	MaxOpenConnections int           `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int           `mapstructure:"maxIdleConnections"`
	MaxIdleTime        time.Duration `mapstructure:"maxIdleTime"`
	MaxOpenConns       int           `mapstructure:"maxOpenConns"`
	MaxIdleConns       int           `mapstructure:"maxIdleConns"`
}

type AuthenticationCofnig struct {
	AccessSecretKey       string        `mapstructure:"accessSecretKey"`
	RefreshSecretKey      string        `mapstructure:"refreshSecretKey"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"accessTokenExpiresIn"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"refreshTokenExpiresIn"`
}
