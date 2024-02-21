package main

import (
	"github.com/alireza-fa/phone-book/cmd"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const description = "PhoneBook Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%v", err)
	}
}
