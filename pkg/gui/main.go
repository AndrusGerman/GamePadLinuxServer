package gui

import (
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
	setStatus("Server Close")
	// Create devices
	w, server := create_app()

	// title status
	title := createTitle()

	// devices count
	bindDevicesCount := binding.IntToStringWithFormat(utils.DevicesConnect, "Devices connect: %d")
	devicesConnect := widget.NewLabel("")
	devicesConnect.Bind(bindDevicesCount)

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
		startServerBtn,
		devicesConnect,
	))

	w.ShowAndRun()
}
