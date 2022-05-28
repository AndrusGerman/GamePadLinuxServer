package gui

import (
	"game_pad_linux_server/pkg/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func listDeices() fyne.CanvasObject {
	listDevicesText := widget.NewLabel("Devices USB:")
	listDevicesText.TextStyle.Bold = true

	listDevices := widget.NewListWithData(
		utils.DevicesList,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		})

	return container.NewVBox(
		listDevicesText,
		listDevices)
}
