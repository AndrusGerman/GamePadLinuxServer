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

var primaryLabelText = "GamePadLinux:"
var statusServer = binding.NewString()

func setStatus(status string) {
	statusServer.Set(primaryLabelText + " (" + status + ")")
}

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

	title := widget.NewLabel(primaryLabelText)
	title.Bind(statusServer)

	setStatus("Server Close")

	title.TextStyle.Bold = true

	ipText := widget.NewLabel("IP:// ")

	bindDevicesCount := binding.IntToStringWithFormat(utils.DevicesConnect, "Devices connect: %d")

	devicesConnect := widget.NewLabel("")

	devicesConnect.Bind(bindDevicesCount)

	ips := utils.GetLocalIP()
	listIP := widget.NewList(
		func() int {
			return len(ips)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(ips[i])
		})

	listDevicesText := widget.NewLabel("Devices USB:")

	listDevices := widget.NewListWithData(
		utils.DevicesList,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		})

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

	listContainer := container.NewGridWithColumns(2,
		container.NewVBox(
			ipText,
			listIP),
		container.NewVBox(
			listDevicesText,
			listDevices),
	)

	w.SetContent(container.NewVBox(
		title,

		listContainer,

		startServer,
		devicesConnect,
	))

	w.ShowAndRun()
}
