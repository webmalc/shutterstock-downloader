package logger

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	IsDebug bool
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		IsDebug: !viper.GetBool("is_prod"),
	}

	return config
}
