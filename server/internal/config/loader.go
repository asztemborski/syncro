package config

import (
	"fmt"
	"os"

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

	if l.envPrefix != "" {
		l.viper.SetEnvPrefix(l.envPrefix)
	}

	if err := l.loadConfigFile(l.configFile, false); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := l.loadConfigFile(l.envFile, true); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error reading .env file: %w", err)
		}
	}

	var config Configuration
	if err := l.viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

func (l *Loader) loadConfigFile(filePath string, merge bool) error {
	if filePath == "" {
		return nil
	}

	l.viper.SetConfigFile(filePath)
	if merge {
		return l.viper.MergeInConfig()
	}

	return l.viper.ReadInConfig()
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
