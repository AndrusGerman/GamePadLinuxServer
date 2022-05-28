package gui

import (
	"game_pad_linux_server/pkg/adb"
	"game_pad_linux_server/pkg/devices"
	"game_pad_linux_server/pkg/server"
	"game_pad_linux_server/pkg/utils"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func Execute() {

	// Create devices
	devices, err := devices.CreateDevices()
	if err != nil {
		log.Println("error start devices: ", err)
		os.Exit(1)
		return
	}
	defer devices.Close()

	// Start Server
	server := server.NewServer(devices)
	defer server.Close()
	go adb.WaitADBClients()

	a := app.NewWithID("com.andruscodex.gamepadlinux")
	w := a.NewWindow("GamePadLinux: AndrusCodex")

	w.Resize(fyne.NewSize(400, 200))

	title := widget.NewLabel("GamePadLinux: (Client GUI)")

	bindDevicesCount := binding.IntToStringWithFormat(utils.DevicesConnect, "Devices connect: %d")

	devicesConnect := widget.NewLabel("")

	devicesConnect.Bind(bindDevicesCount)

	ips := utils.GetLocalIP()
	list := widget.NewList(
		func() int {
			return len(ips)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ips[i])
		})

	startServer := widget.NewButton("Start Server", func() {
		server.Close()
		go func() {
			err := server.Server("8992")
			if err != nil {
				log.Println("error start server: ", err)
				return
			}
		}()
	})

	w.SetContent(container.NewVBox(
		title,
		list,
		devicesConnect,
		startServer,
	))

	w.ShowAndRun()
}
