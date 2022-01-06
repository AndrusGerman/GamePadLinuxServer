package main

import (
	"github.com/bendahl/uinput"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(mouse uinput.Mouse, keyboard uinput.Keyboard) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/ws", handlerEvents)

	e.GET("/storage/get/:name", StorageHandlerGet)
	e.POST("/storage/set/:name", StorageHandlerSet)

	activate_events(mouse, keyboard)
	go proccess_events()
	e.Logger.Fatal(e.Start(":1323"))
}
