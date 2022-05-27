package app

import (
	"strings"

	"github.com/bendahl/uinput"
)

func (ctx *Events) ManagerKeybord(keyboard uinput.Keyboard) {
	switch ctx.Mode {
	case 3:
		vl, ok := keyMap[ctx.Value]
		if ok {
			keyboard.KeyUp(int(vl))
		}
	case 2:
		ctx.Value = strings.Title(ctx.Value)
		vl, ok := keyMap[ctx.Value]
		if ok {
			keyboard.KeyPress(int(vl))
		}
	case 1:
		vl, ok := keyMap[ctx.Value]
		if ok {
			keyboard.KeyDown(int(vl))
		}
	}
}
