package cmd

import (
	"github.com/alireza-fa/phone-book/internal/config"
	"github.com/alireza-fa/phone-book/internal/models"
	"github.com/alireza-fa/phone-book/internal/repository"
	"github.com/alireza-fa/phone-book/pkg/logger"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type Migrate struct{}

func (m Migrate) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.run(config.Load(true), args, trap)
	}

	return &cobra.Command{
		Use:       "migrate",
		Short:     "run migrations",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m Migrate) run(cfg *config.Config, args []string, trap chan os.Signal) {
	log := logger.NewZap(cfg.Logger)

	if len(args) != 1 {
		log.Fatal("Invalid arguments given", zap.Any("args", args))
	}

	rd, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		log.Panic("Error creating rdbms", zap.Error(err))
	}

	repo := repository.New(log, cfg.Repository, rd)
	if err := repo.Migrate(models.Migrate(args[0])); err != nil {
		log.Fatal("Error while migrating", zap.String("direction", args[0]), zap.Error(err))
	}

	log.Info("Database has been migrated successfully", zap.String("direction", args[0]))
}
