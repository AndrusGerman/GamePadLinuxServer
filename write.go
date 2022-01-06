package main

import (
	"strings"
	"unicode"

	"github.com/bendahl/uinput"
	"github.com/go-vgo/robotgo"
)

func (ctx *Events) ManagerWriter(keyboard uinput.Keyboard) {
	if strings.Contains(ctx.Value, "{") && strings.Contains(ctx.Value, "}") {
		switch ctx.Value {
		case "{bksp}":
			keyboard.KeyPress(14)
		case "{enter}":
			keyboard.KeyPress(28)
		case "{space}":
			keyboard.KeyPress(57)
		case "{tab}":
			keyboard.KeyPress(15)
		}
		return
	}

	if value, ok := keyMap[strings.ToUpper(ctx.Value)]; ok {
		if unicode.IsLetter(rune(ctx.Value[0])) && strings.ToUpper(ctx.Value) == ctx.Value {
			keyboard.KeyDown(uinput.KeyLeftshift)
			keyboard.KeyPress(int(value))
			keyboard.KeyUp(uinput.KeyLeftshift)
		} else {
			keyboard.KeyPress(int(value))
		}

	} else {
		robotgo.TypeStr(ctx.Value)
	}
}
