package devices

import (
	"log"

	"github.com/bendahl/uinput"
)

type Devices interface {
	GetKeyboard() uinput.Keyboard
	GetMouse() uinput.Mouse
}

type DefaultDevices struct {
	keyboard uinput.Keyboard
	mouse    uinput.Mouse
}

func (ctx *DefaultDevices) Close() error {
	if ctx.keyboard != nil {
		ctx.keyboard.Close()
	}
	if ctx.mouse != nil {
		ctx.mouse.Close()
	}
	return nil
}

func (ctx *DefaultDevices) GetKeyboard() uinput.Keyboard {
	return ctx.keyboard
}
func (ctx *DefaultDevices) GetMouse() uinput.Mouse {
	return ctx.mouse
}

func (ctx *DefaultDevices) StartDevice() error {
	ctx.Close()
	var err error
	// Create devices
	ctx.keyboard, err = uinput.CreateKeyboard("/dev/uinput", []byte("AndrusCodex / Keyborde"))
	if err != nil {
		log.Println("Keybord: ", err)
		return err
	}
	ctx.mouse, err = uinput.CreateMouse("/dev/uinput", []byte("AndrusCodex / Mouse"))
	if err != nil {
		defer ctx.keyboard.Close()
		log.Println("Mouse: ", err)
		return err
	}
	return nil
}

func CreateDevices() (*DefaultDevices, error) {
	var devices = new(DefaultDevices)
	// Create devices
	return devices, devices.StartDevice()

}
