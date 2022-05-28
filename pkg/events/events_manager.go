package events

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/bendahl/uinput"
)

type EventsManagerDefault struct {
	// Events WS
	EnventsChan chan *ManagerWS

	// Manager all Chanels
	ManagerMouseChan            chan *Events
	ManagerKeybordChan          chan *Events
	ManagerKeyMouseChan         chan *Events
	ManagerLatencyChan          chan *ManagerWS
	ManagerJoystickKeyboardChan chan *Events
	ManagerWriterChan           chan *Events

	// Devices
	mouse    uinput.Mouse
	keyboard uinput.Keyboard
	//  Manager Devices
	devices devices.Devices
}

type EventsMangar interface {
	Close() error
	GetEnventsChan() chan *ManagerWS
}

func NewEventsManager(devices devices.Devices) EventsMangar {
	var em = new(EventsManagerDefault)
	em.StartChanels()
	em.SetDevices(devices)
	em.ActivateEvents()
	go em.ProccessEvents()
	return em
}
