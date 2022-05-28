package server

import (
	"encoding/json"
	"fmt"
	"game_pad_linux_server/pkg/adb"
	"game_pad_linux_server/pkg/events"
	"log"

	"github.com/labstack/echo/v4"
)

func (ctx *ServerManagerDefault) handlerEvents(c echo.Context) error {
	fmt.Println("Se conecto un cliente")
	adb.DevicesConnect++
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	defer fmt.Println("Salio cliente")
	defer func() {
		adb.DevicesConnect--
	}()

	// events
	evmanager := events.NewEventsManager(ctx.devices)
	defer evmanager.Close()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			//ws.RemoteAddr().Network()
			log.Println("Error Read Message, close: ", err)
			return nil
		}
		go func() {
			var ev = new(events.Events)
			err = json.Unmarshal(msg, ev)
			if err != nil {
				log.Println("Error Unmarshal: ", err)
				return
			}

			// Send Event
			evmanager.GetEnventsChan() <- &events.ManagerWS{
				WS:     ws,
				Events: ev,
			}
		}()
	}

}