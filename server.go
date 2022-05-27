package main

import (
	"game_pad_linux_server/pkg/devices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(devices devices.Devices) {
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/ws", handlerEvents)

	e.GET("/storage/get/:name", StorageHandlerGet)
	e.POST("/storage/set/:name", StorageHandlerSet)
	e.GET("/open", EnabledHandlerGet)

	e.Logger.Fatal(e.Start(":8992"))
}
