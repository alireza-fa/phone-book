package config

import (
	"github.com/alireza-fa/phone-book/internal/repository"
	"github.com/alireza-fa/phone-book/pkg/logger"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"github.com/alireza-fa/phone-book/pkg/token"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	Token      *token.Config      `koanf:"token"`
	RDBMS      *rdbms.Config      `koanf:"rdbms"`
	Repository *repository.Config `koanf:"repository"`
}
