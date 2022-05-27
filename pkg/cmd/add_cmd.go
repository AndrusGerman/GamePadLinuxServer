package cmd

import (
	"game_pad_linux_server/pkg/utils"
)

var useGui bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&useGui, "gui", "g", utils.IsGUI(), "start 'gamepad' in gui mode")
}
