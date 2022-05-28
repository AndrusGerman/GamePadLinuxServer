package gui

import (
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/server"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func create_app() (fyne.Window, *server.ServerManagerDefault) {
	// Create devices
	devices, err := devices.CreateDevices()
	if err != nil {
		log.Println("error start devices: ", err)
		os.Exit(1)
	}
	defer devices.Close()

	// Start Server
	server := server.NewServer(devices)
	defer server.Close()

	a := app.NewWithID("com.andruscodex.gamepadlinux")
	w := a.NewWindow("GamePadLinux: AndrusCodex")

	w.Resize(fyne.NewSize(400, 200))
	return w, server
}
