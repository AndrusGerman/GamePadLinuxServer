package events

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

// To measure the latency you send the last value of time now, which the server answered you
// and in this way the first value you will get is the current latency and the new value of time now
// you send -> Last time.Now() value ($2)  | from the last reply     $1|$2
// --
// you get <- $1|$2 | example 20|232323232323
// the first value is the latency, and the second is the last time.Now()
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

// -> 0
// <- Any|200000
// -> 200000
// <- 3|201003 // your latency is 3
// -> 201003
// <- 20|202023 // your latency is 20...
