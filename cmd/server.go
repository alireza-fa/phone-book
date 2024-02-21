package cmd

import (
	"github.com/alireza-fa/phone-book/internal/config"
	"github.com/alireza-fa/phone-book/pkg/logger"
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

	log.Error("IMPLEMENT ME !!!!")

	filed := zap.String("signal trap", (<-trap).String())
	log.Info("exiting by receiving unix signal", filed)
}
