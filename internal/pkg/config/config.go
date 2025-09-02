package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application.
type Config struct {
	Server struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"server"`
}

// Load reads configuration from a file.
func Load(appName string) (*Config, error) {
	v := viper.New()
	v.SetConfigName(appName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
