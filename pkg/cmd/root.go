package cmd

import (
	"fmt"
	"game_pad_linux_server/app"
	"game_pad_linux_server/pkg/gui"
	"os"

	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "game_pad_linux_server",
	Short: "server to create wireless controllers for linux",
	Run: func(cmd *cobra.Command, args []string) {
		welcome()
		if useGui {
			gui.Execute()
		}
		if !useGui {
			app.Execute()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func welcome() {
	ms1 := color.Green("GamePad: by AndrusCodex is start 🔥")
	ms2 := color.Yellow("GamePad: Please report any errors in")
	ms3 := color.Yellow("https://github.com/AndrusGerman/GamePadLinuxServer")
	ms4 := color.Grey("GamePad: Your contribution can make big changes")
	fmt.Printf("%s\n%s\n%s\n%s\n", ms1, ms2, ms3, ms4)
}
