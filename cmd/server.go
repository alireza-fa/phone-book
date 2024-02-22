package cmd

import (
	"github.com/alireza-fa/phone-book/internal/api/http"
	"github.com/alireza-fa/phone-book/internal/config"
	"github.com/alireza-fa/phone-book/internal/repository"
	"github.com/alireza-fa/phone-book/pkg/logger"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"github.com/alireza-fa/phone-book/pkg/token"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run PhoneBook server",
		Run:   run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {
	log := logger.NewZap(cfg.Logger)

	t, err := token.New(cfg.Token)
	if err != nil {
		log.Panic("Error creating token object", zap.Error(err))
	}

	rd, err := rdbms.New(cfg.RDBMS)
	if err != nil {
		log.Panic("error creating rdbms", zap.Error(err))
	}

	repo := repository.New(log, cfg.Repository, rd)

	server := http.New(log, t, repo)
	go server.Serve(8080)

	filed := zap.String("signal trap", (<-trap).String())
	log.Info("exiting by receiving unix signal", filed)
}
