package app

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(port string, devices devices.Devices) error {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// server route
	e.GET("/ws", handlerEvents)
	e.GET("/open", EnabledHandlerGet)

	return e.Start(":" + port)
}

func EnabledHandlerGet(ctx echo.Context) error {
	return ctx.String(200, "Is Open")
}
