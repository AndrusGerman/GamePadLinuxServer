package utils

import "fyne.io/fyne/v2/data/binding"

var StatusLog = binding.NewString()

func SetStatusLog(message string) {
	StatusLog.Set("Log: " + message)
}
