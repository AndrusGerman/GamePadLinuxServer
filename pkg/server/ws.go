package server

import (
	"encoding/json"
	"fmt"
	"game_pad_linux_server/pkg/events"
	"game_pad_linux_server/pkg/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

func (ctx *ServerManagerDefault) handlerEvents(c echo.Context) error {
	fmt.Println(color.Grey("gamepad-server: A device connected"))
	utils.SumDevicesConnect(1)
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	defer fmt.Println(color.Grey("gamepad-server: A device disconnected"))
	defer utils.SumDevicesConnect(-1)

	// events
	evmanager := events.NewEventsManager(ctx.devices)
	defer evmanager.Close()

	var event = evmanager.GetEnventsChan()

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
			event <- &events.ManagerWS{
				WS:     ws,
				Events: ev,
			}
		}()
	}

}
