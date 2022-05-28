package gui

import "fyne.io/fyne/v2/widget"

func createTitle() *widget.Label {
	title := widget.NewLabel(primaryLabelText)
	title.Bind(statusServer)
	title.TextStyle.Bold = true
	title.TextStyle.Monospace = true

	return title
}
