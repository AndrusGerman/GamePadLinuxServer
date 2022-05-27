package cmd

import (
	"os"

	"github.com/mattn/go-isatty"
)

var useGui bool

func init() {

	rootCmd.PersistentFlags().BoolVarP(&useGui, "gui", "g", useGuiDefault(), "start 'gamepad' in gui mode")
}

func useGuiDefault() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return true
	}
	return false
}
