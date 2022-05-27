package app

import (
	"fmt"
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
)

func Server(port string, devices devices.Devices) error {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// server route
	e.GET("/ws", handlerEvents)
	e.GET("/open", EnabledHandlerGet)
	fmt.Println(color.Bold("GamePad: server start on port "), color.Green(":"+port))
	fmt.Println(color.Bold("GamePad: server start ip: "), color.Green(utils.GetLocalIP()))
	return e.Start(":" + port)
}

func EnabledHandlerGet(ctx echo.Context) error {
	return ctx.String(200, "Is Open")
}
