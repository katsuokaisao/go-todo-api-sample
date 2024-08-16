package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/katsuokaisao/gin-play/domain"
	"github.com/spf13/cobra"
)

var jwtCreateCmd = &cobra.Command{
	Use:   "jwt-create",
	Short: "Create a JWT",
	Run: func(cmd *cobra.Command, args []string) {
		userUUID := uuid.New().String()
		scope, ok := domain.ToScope(jwtCreateCmdScopeFlag)
		if !ok {
			panic(fmt.Errorf("invalid scope: %s", jwtCreateCmdScopeFlag))
		}
		scopes := []domain.Scope{scope}

		e := domain.NewEnv()
		if err := e.Load(); err != nil {
			panic(err)
		}
		g := domain.NewJWTGenerator(e.API.JwtSecret)
		token, err := g.Generate(userUUID, scopes)
		if err != nil {
			panic(fmt.Errorf("failed to generate jwt: %w", err))
		}
		fmt.Printf("userUUID: %s\n", userUUID)
		fmt.Printf("scope: %s\n", scope)
		fmt.Printf("token: %s\n", token)
	},
}
