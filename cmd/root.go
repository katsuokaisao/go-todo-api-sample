package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(serveCmd)
}

var rootCmd = &cobra.Command{
	Use: "",
}

func Execute() {
	rootCmd.Execute()
}
