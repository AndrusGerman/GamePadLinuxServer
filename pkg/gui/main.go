package gui

import (
	"game_pad_linux_server/pkg/adb"
	"game_pad_linux_server/pkg/utils"

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

	// adb server
	watch := adb.WaitADBClients()
	defer watch.Close()

	setStatus("Server Close")
	// Create devices
	w, server, devices := create_app()
	defer devices.Close()

	// title status
	title := createTitle()

	// devices count
	bindDevicesCount := binding.IntToStringWithFormat(utils.DevicesConnect, "Devices connect: %d")
	devicesConnect := widget.NewLabel("")
	devicesConnect.TextStyle.Bold = true
	devicesConnect.Bind(bindDevicesCount)

	// devices count
	statusLogs := widget.NewLabel("")
	statusLogs.TextStyle.Italic = true
	statusLogs.Bind(utils.StatusLog)

	// create server btn
	var startServerBtn = createStartServerBtn(server)

	// list container
	listContainer := container.NewGridWithColumns(2,
		listDevices(),
		listDeices(),
	)

	w.SetContent(container.NewVBox(
		title,
		listContainer,
		devicesConnect,
		statusLogs,
		startServerBtn,
	))

	w.ShowAndRun()
}
