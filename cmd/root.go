package cmd

import "github.com/spf13/cobra"

var (
	jwtCreateCmdScopeFlag string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(jwtCreateCmd)

	jwtCreateCmd.Flags().StringVarP(&jwtCreateCmdScopeFlag, "scope", "s", "", "scope")
}

var rootCmd = &cobra.Command{
	Use: "",
}

func Execute() {
	rootCmd.Execute()
}
