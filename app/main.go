package app

import (
	"fmt"
	"game_pad_linux_server/pkg/adb"
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/server"
	"log"
	"os"

	"github.com/labstack/gommon/color"
)

func Execute() {
	fmt.Println(color.Magenta("gamepad-cli: please test 'game_pad_linux_server --gui'"))
	// Create devices
	devices, err := devices.CreateDevices()
	if err != nil {
		log.Println("error start devices: ", err)
		os.Exit(1)
		return
	}
	defer devices.Close()

	watch := adb.WaitADBClients()
	defer watch.Close()

	// Start Server
	server := server.NewServer(devices)
	err = server.Server("8992")
	if err != nil {
		log.Println("error start server: ", err)
		os.Exit(1)
		return
	}
}
