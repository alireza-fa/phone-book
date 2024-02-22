package config

import (
	"github.com/alireza-fa/phone-book/pkg/logger"
	"github.com/alireza-fa/phone-book/pkg/token"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	Token  *token.Config  `koanf:"token"`
}
