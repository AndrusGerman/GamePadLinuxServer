package app

import (
	"game_pad_linux_server/pkg/adb"
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/events"
	"log"
	"os"
)

func Execute() {

	// Create devices
	devices, err := devices.CreateDevices()
	if err != nil {
		log.Println("error start devices: ", err)
		os.Exit(1)
		return
	}
	defer devices.Close()

	go adb.WaitADBClients()

	// events
	events.ActivateEvents(devices)
	go events.ProccessEvents()

	// Start Server
	err = Server("8992", devices)
	if err != nil {
		log.Println("error start server: ", err)
		os.Exit(1)
		return
	}
}
