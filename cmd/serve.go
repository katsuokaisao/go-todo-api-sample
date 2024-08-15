package cmd

import (
	"fmt"

	"github.com/katsuokaisao/gin-play/domain"
	"github.com/katsuokaisao/gin-play/wire/serve"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		e := domain.NewEnv()
		if err := e.Load(); err != nil {
			panic(err)
		}
		serve, err := serve.SetUpServe(e)
		if err != nil {
			panic(fmt.Errorf("failed to set up serve: %w", err))
		}
		serve.Server.RegisterRoutes()
		serve.Server.Run()
	},
}
