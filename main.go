package main

import (
	"game_pad_linux_server/pkg/devices"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create devices
	devices, err := devices.CreateDevices()
	if err != nil {
		log.Println("error start devices: ", err)
		return
	}
	defer devices.Close()

	go waitADBClients()

	// events
	activate_events(devices)
	go proccess_events()

	// Start Server
	Server(devices)

}
