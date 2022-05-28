package events

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/gorilla/websocket"
)

type ManagerWS struct {
	*Events
	WS *websocket.Conn
}

const (
	TypeManagerMouseChan            = 1
	TypeManagerKeybordChan          = 2
	TypeManagerKeyMouseChan         = 3
	TypeManagerLatencyChan          = 5
	TypeManagerJoystickKeyboardChan = 6
	TypeManagerWriterChan           = 7
)

type Events struct {
	Type   uint
	Value  string
	ValueX float32
	ValueY float32
	Mode   uint
}

func (ctx *EventsManagerDefault) StartChanels() {
	// Events WS
	ctx.EnventsChan = make(chan *ManagerWS, 10)

	// Manager all Chanels
	ctx.ManagerMouseChan = make(chan *Events, 10)
	ctx.ManagerKeybordChan = make(chan *Events, 10)
	ctx.ManagerKeyMouseChan = make(chan *Events, 10)
	ctx.ManagerLatencyChan = make(chan *ManagerWS, 10)
	ctx.ManagerJoystickKeyboardChan = make(chan *Events, 10)
	ctx.ManagerWriterChan = make(chan *Events, 10)
}

func (ctx *EventsManagerDefault) CloseChanels() {
	// Events WS
	close(ctx.EnventsChan)

	// Manager all Chanels
	close(ctx.ManagerMouseChan)
	close(ctx.ManagerKeybordChan)
	close(ctx.ManagerKeyMouseChan)
	close(ctx.ManagerLatencyChan)
	close(ctx.ManagerJoystickKeyboardChan)
	close(ctx.ManagerWriterChan)
}

func (ctx *EventsManagerDefault) Close() error {
	ctx.CloseChanels()
	return nil
}

// send the events, to the device's event handler
func (ctx *EventsManagerDefault) ProccessEvents() {
	go func() {
		for ev := range ctx.EnventsChan {
			switch ev.Type {
			case TypeManagerMouseChan:
				ctx.ManagerMouseChan <- ev.Events
			case TypeManagerKeybordChan:
				ctx.ManagerKeybordChan <- ev.Events
			case TypeManagerKeyMouseChan:
				ctx.ManagerKeyMouseChan <- ev.Events
			case TypeManagerLatencyChan:
				ctx.ManagerLatencyChan <- ev
			case TypeManagerJoystickKeyboardChan:
				ctx.ManagerJoystickKeyboardChan <- ev.Events
			case TypeManagerWriterChan:
				ctx.ManagerWriterChan <- ev.Events
			}
		}
	}()
}

func (ctx *EventsManagerDefault) SetDevices(devices devices.Devices) {
	ctx.devices = devices
	ctx.mouse = devices.GetMouse()
	ctx.keyboard = devices.GetKeyboard()
}

// activate the devices event handler
func (ctx *EventsManagerDefault) ActivateEvents() {

	// Get Mouse Events
	go func() {
		for ev := range ctx.ManagerMouseChan {
			ev.ManagerMouse(ctx.mouse)
		}
	}()

	// get Keyboard Events
	go func() {
		for ev := range ctx.ManagerKeybordChan {
			ev.ManagerKeybord(ctx.keyboard)
		}
	}()

	// Get Mouse Clicks
	go func() {
		for ev := range ctx.ManagerKeyMouseChan {
			ev.ManagerKeyMouse(ctx.mouse)
		}
	}()

	// Get Latency Events
	go func() {
		for ev := range ctx.ManagerLatencyChan {
			ev.ManagerLatency(ev.WS)
		}
	}()

	// Get Joystick to Keyboard
	go func() {
		for ev := range ctx.ManagerJoystickKeyboardChan {
			ev.ManagerJoystickKeyboard(ctx.keyboard)
		}
	}()

	// Get Writers events
	go func() {
		for ev := range ctx.ManagerWriterChan {
			ev.ManagerWriter(ctx.keyboard)
		}
	}()
}

func (ctx *EventsManagerDefault) GetEnventsChan() chan *ManagerWS {
	return ctx.EnventsChan
}
