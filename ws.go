package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bendahl/uinput"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

func handlerEvents(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error Read Message, close: ", ws.RemoteAddr().Network())
			return nil
		}
		var ev = new(Events)
		err = json.Unmarshal(msg, ev)
		if err != nil {
			log.Println("Error Unmarshal: ", ws.RemoteAddr().Network())
			continue
		}
		// Send Event
		enventsChan <- &ManagerWS{
			ws:     ws,
			Events: ev,
		}
	}

}

type Events struct {
	Type   uint
	Value  string
	ValueX float32
	ValueY float32
	Mode   uint
}

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

func (ctx *Events) ManagerLatency(ws *websocket.Conn) {
	switch ctx.Mode {
	case 1:
		ws.WriteMessage(websocket.TextMessage, []byte("0|0"))
	case 2:
		n, _ := strconv.Atoi(ctx.Value)
		var now = time.Now().UnixNano()
		var tt = now - int64(n)
		time.Sleep(1000 * time.Millisecond)
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d|%d", tt, time.Now().UnixNano())))
	}
}

var pxMovementMax = 12

func (ctx *Events) ManagerMouse(mouse uinput.Mouse) {
	switch ctx.Mode {
	case 3:
		pxEvent = nil
	case 1:
		if move {
			move = false
			clearMove()
		}
		universalMouse = mouse
		pxEvent = ctx
	}
}

func (ctx *Events) UsedMouse(mouse uinput.Mouse) {
	switch ctx.Mode {
	case 3:
		pxEvent = nil
	case 1:
		mouse.Move(
			int32(math.Round(float64(pxMovementMax)*float64(ctx.ValueX))),
			int32(math.Round(float64(pxMovementMax)*float64(ctx.ValueY))),
		)
	}
}

var pxEvent *Events
var move = true
var universalMouse uinput.Mouse

func clearMove() {
	go func() {
		time.Sleep(6 * time.Millisecond)
		move = true
	}()
}

func init() {
	go func() {
		for {
			if universalMouse != nil {
				if pxEvent != nil {
					pxEvent.UsedMouse(universalMouse)
				}
			}
			time.Sleep(16 * time.Millisecond)
		}
	}()
}

func (ctx *Events) ManagerJoystickKeyboard(keyboard uinput.Keyboard) {
	switch ctx.Mode {
	case 1:
		keys := strings.Split(ctx.Value, ",")
		for _, v := range keys {
			// Press Buttom
			key := strings.Split(v, "-")

			// Preciona la selecionada
			vl, ok := keyMap[key[0]]
			if ok {
				keyboard.KeyDown(int(vl))
			}

			// Suelta la contraria
			vl, ok = keyMap[key[1]]
			if ok {
				keyboard.KeyUp(int(vl))
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
