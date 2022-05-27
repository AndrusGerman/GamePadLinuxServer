package cmd

import (
	"game_pad_linux_server/app"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "game_pad_linux_server",
	Short: "server to create wireless controllers for linux",
	Run: func(cmd *cobra.Command, args []string) {
		app.Execute()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
