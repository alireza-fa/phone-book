package config

import "github.com/alireza-fa/phone-book/pkg/logger"

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
	}
}
