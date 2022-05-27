package app

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/gorilla/websocket"
)

type ManagerWS struct {
	*Events
	ws *websocket.Conn
}

// Events WS
var enventsChan = make(chan *ManagerWS, 10)

// Manager all Chanels
var ManagerMouseChan = make(chan *Events, 10)
var ManagerKeybordChan = make(chan *Events, 10)
var ManagerKeyMouseChan = make(chan *Events, 10)
var ManagerLatencyChan = make(chan *ManagerWS, 10)
var ManagerJoystickKeyboardChan = make(chan *Events, 10)
var ManagerWriterChan = make(chan *Events, 10)

const (
	TypeManagerMouseChan            = 1
	TypeManagerKeybordChan          = 2
	TypeManagerKeyMouseChan         = 3
	TypeManagerLatencyChan          = 5
	TypeManagerJoystickKeyboardChan = 6
	TypeManagerWriterChan           = 7
)

func ProccessEvents() {
	go func() {
		for ev := range enventsChan {
			switch ev.Type {
			case TypeManagerMouseChan:
				ManagerMouseChan <- ev.Events
			case TypeManagerKeybordChan:
				ManagerKeybordChan <- ev.Events
			case TypeManagerKeyMouseChan:
				ManagerKeyMouseChan <- ev.Events
			case TypeManagerLatencyChan:
				ManagerLatencyChan <- ev
			case TypeManagerJoystickKeyboardChan:
				ManagerJoystickKeyboardChan <- ev.Events
			case TypeManagerWriterChan:
				ManagerWriterChan <- ev.Events
			}
		}
	}()
}

func ActivateEvents(devices devices.Devices) {
	var mouse = devices.GetMouse()
	var keyboard = devices.GetKeyboard()

	// Get Mouse Events
	go func() {
		for ev := range ManagerMouseChan {
			ev.ManagerMouse(mouse)
		}
	}()

	// get Keyboard Events
	go func() {
		for ev := range ManagerKeybordChan {
			ev.ManagerKeybord(keyboard)
		}
	}()

	// Get Mouse Clicks
	go func() {
		for ev := range ManagerKeyMouseChan {
			ev.ManagerKeyMouse(mouse)
		}
	}()

	// Get Latency Events
	go func() {
		for ev := range ManagerLatencyChan {
			ev.ManagerLatency(ev.ws)
		}
	}()

	// Get Joystick to Keyboard
	go func() {
		for ev := range ManagerJoystickKeyboardChan {
			ev.ManagerJoystickKeyboard(keyboard)
		}
	}()

	// Get Writers events

	go func() {
		for ev := range ManagerWriterChan {
			ev.ManagerWriter(keyboard)
		}
	}()
}

// func close_chanels() {
// 	close(enventsChan)
// 	close(ManagerMouseChan)
// 	close(ManagerKeybordChan)
// 	close(ManagerKeyMouseChan)
// 	close(ManagerLatencyChan)
// 	close(ManagerJoystickKeyboardChan)
// 	close(ManagerWriterChan)
// }

// func start_chanels() {
// 	// Events WS
// 	enventsChan = make(chan *ManagerWS, 10)

// 	// Manager all Chanels
// 	ManagerMouseChan = make(chan *Events, 10)
// 	ManagerKeybordChan = make(chan *Events, 10)
// 	ManagerKeyMouseChan = make(chan *Events, 10)
// 	ManagerLatencyChan = make(chan *ManagerWS, 10)
// 	ManagerJoystickKeyboardChan = make(chan *Events, 10)
// 	ManagerWriterChan = make(chan *Events, 10)

// }
