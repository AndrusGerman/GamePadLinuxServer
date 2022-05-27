package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

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
