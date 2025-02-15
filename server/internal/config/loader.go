package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Option func(*Loader)

type ConfigLoader interface {
	Load() (*Configuration, error)
}

type Loader struct {
	viper      *viper.Viper
	configFile string
	envFile    string
	envPrefix  string
	defaults   map[string]any
}

func NewLoader(options ...Option) *Loader {
	loader := &Loader{
		viper: viper.New(),
	}

	for _, option := range options {
		option(loader)
	}

	return loader
}

func (l *Loader) Load() (*Configuration, error) {
	for key, value := range l.defaults {
		l.viper.SetDefault(key, value)
	}

	if l.configFile != "" {
		l.viper.SetConfigFile(l.configFile)
		if err := l.viper.ReadInConfig(); err != nil {
			if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
				return nil, fmt.Errorf("error reading config file: %w", err)
			}
		}
	}

	if l.envFile != "" {
		if err := godotenv.Load(l.envFile); err != nil {
			if !os.IsNotExist(err) {
				return nil, fmt.Errorf("error loading .env file: %w", err)
			}
		}
	}
	l.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if l.envPrefix != "" {
		l.viper.SetEnvPrefix(l.envPrefix)
	}
	l.viper.AllowEmptyEnv(true)
	l.setEnvBindings(l.viper, l.envPrefix, DefaultConfig)
	l.viper.AutomaticEnv()

	var config Configuration
	if err := l.viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

func (l *Loader) setEnvBindings(v *viper.Viper, prefix string, cfg map[string]any) {
	for key, value := range cfg {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		if nested, ok := value.(map[string]any); ok {
			l.setEnvBindings(v, fullKey, nested)
		} else {
			v.BindEnv(fullKey)
		}
	}
}

func WithConfigFile(configFile string) Option {
	return func(l *Loader) {
		l.configFile = configFile
	}
}

func WithEnvFile(envFile string) Option {
	return func(l *Loader) {
		l.envFile = envFile
	}
}

func WithDefaults(defaults map[string]any) Option {
	return func(l *Loader) {
		l.defaults = defaults
	}
}

func WithEnvPrefix(prefix string) Option {
	return func(l *Loader) {
		l.envPrefix = prefix
	}
}
