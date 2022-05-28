package gui

import (
	"game_pad_linux_server/pkg/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func listDevices() fyne.CanvasObject {
	ipText := widget.NewLabel("IP:// ")
	ipText.TextStyle.Bold = true

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

	return container.NewVBox(
		ipText,
		listIP)
}
