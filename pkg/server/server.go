package server

import (
	"fmt"
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
)

type ServerManagerDefault struct {
	devices devices.Devices
	server  *echo.Echo
}

func (ctx *ServerManagerDefault) Server(port string) error {

	ctx.server = echo.New()
	ctx.server.HideBanner = true
	ctx.server.HidePort = true
	ctx.server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// server route
	ctx.server.GET("/ws", ctx.handlerEvents)
	ctx.server.GET("/open", ctx.enabledHandlerGet)
	fmt.Println(color.Bold("GamePad: server start on port "), color.Green(":"+port))
	fmt.Println(color.Bold("GamePad: server start ip: "), color.Green(utils.GetLocalIP()))
	return ctx.server.Start(":" + port)
}

func (*ServerManagerDefault) enabledHandlerGet(ctx echo.Context) error {
	return ctx.String(200, "Is Open")
}

func (ctx *ServerManagerDefault) Close() error {
	return ctx.server.Close()
}
