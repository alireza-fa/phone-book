package repository

import (
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"go.uber.org/zap"
)

type Repository interface{}

type repository struct {
	logger *zap.Logger
	config *Config
	rdbms  rdbms.RDBMS
}

func New(lg *zap.Logger, cfg *Config, rdbms rdbms.RDBMS) Repository {
	return &repository{
		logger: lg,
		config: cfg,
		rdbms:  rdbms,
	}
}
