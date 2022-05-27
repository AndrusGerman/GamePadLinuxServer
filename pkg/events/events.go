package events

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/gorilla/websocket"
)

type ManagerWS struct {
	*Events
	WS *websocket.Conn
}

// Events WS
var EnventsChan = make(chan *ManagerWS, 10)

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

// send the events, to the device's event handler
func ProccessEvents() {
	go func() {
		for ev := range EnventsChan {
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

// activate the devices event handler
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
			ev.ManagerLatency(ev.WS)
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

type Events struct {
	Type   uint
	Value  string
	ValueX float32
	ValueY float32
	Mode   uint
}
