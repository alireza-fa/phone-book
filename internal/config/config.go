package config

import "github.com/alireza-fa/phone-book/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
