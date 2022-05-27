package utils

import (
	"os"

	"github.com/mattn/go-isatty"
)

func IsGUI() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return false
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return false
	}
	return true
}
