package main

import (
	"log"
	"runtime"

	"github.com/bendahl/uinput"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create devices
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("AndrusCodex / Keyborde"))
	if err != nil {
		log.Println("Keybord: ", err)
		return
	}
	defer keyboard.Close()

	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("AndrusCodex / Mouse"))
	if err != nil {
		log.Println("Mouse: ", err)
		return
	}

	defer mouse.Close()

	// Start Server
	Server(mouse, keyboard)

}
