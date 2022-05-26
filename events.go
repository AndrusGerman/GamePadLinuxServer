package main

import (
	"github.com/bendahl/uinput"
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

func proccess_events() {
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

func activate_events(mouse uinput.Mouse, keyboard uinput.Keyboard) {
	go func() {
		for ev := range ManagerMouseChan {
			ev.ManagerMouse(mouse)
		}
	}()

	go func() {
		for ev := range ManagerKeybordChan {
			ev.ManagerKeybord(keyboard)
		}
	}()
	go func() {
		for ev := range ManagerKeyMouseChan {
			ev.ManagerKeyMouse(mouse)
		}
	}()
	go func() {
		for ev := range ManagerLatencyChan {
			ev.ManagerLatency(ev.ws)
		}
	}()
	go func() {
		for ev := range ManagerJoystickKeyboardChan {
			ev.ManagerJoystickKeyboard(keyboard)
		}
	}()
	go func() {
		for ev := range ManagerWriterChan {
			ev.ManagerWriter(keyboard)
		}
	}()
}
