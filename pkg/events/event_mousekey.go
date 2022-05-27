package events

import "github.com/bendahl/uinput"

func (ctx *Events) ManagerKeyMouse(mouse uinput.Mouse) {
	switch ctx.Mode {
	case 3:
		switch ctx.Value {
		case "L":
			mouse.LeftRelease()
		case "R":
			mouse.RightRelease()
		}
	case 2:
		switch ctx.Value {
		case "L":
			mouse.LeftClick()
		case "R":
			mouse.RightClick()
		}
	case 1:
		switch ctx.Value {
		case "L":
			mouse.LeftPress()
		case "R":
			mouse.RightPress()
		}
	}
}
