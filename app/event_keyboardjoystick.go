package app

import (
	"strings"

	"github.com/bendahl/uinput"
)

func (ctx *Events) ManagerJoystickKeyboard(keyboard uinput.Keyboard) {
	switch ctx.Mode {
	case 1:
		keyType := strings.Split(ctx.Value, "|")
		keysAdd := strings.Split(keyType[0], ",")
		keyRemove := strings.Split(keyType[1], ",")
		for _, key := range keyRemove {

			// Suelta las demas
			vl, ok := keyMap[key]
			if ok {
				keyboard.KeyUp(int(vl))
			}

		}

		for _, key := range keysAdd {

			// Preciona la selecionada
			vl, ok := keyMap[key]
			if ok {
				keyboard.KeyDown(int(vl))
			}

		}
	case 3:
		keys := strings.Split(ctx.Value, ",")
		for _, v := range keys {
			// Suelta la tecla
			vl, ok := keyMap[v]
			if ok {
				keyboard.KeyUp(int(vl))
			}
		}
	}
}
