package gui

import (
	"game_pad_linux_server/pkg/server"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func createStartServerBtn(server *server.ServerManagerDefault) fyne.CanvasObject {
	var startServer *widget.Button
	startServer = widget.NewButton("Start Server", func() {
		startServer.SetText("Waiting...")
		startServer.Disable()
		server.Close()
		setStatus("Server Start")
		go func() {
			startServer.Enable()
			startServer.SetText("Restart Server")
			err := server.Server("8992")
			if err != nil {
				setStatus("Server close...")
				log.Println("error start server: ", err)
				return
			}
			startServer.SetText("Start Server")
		}()
	})

	return startServer
}
