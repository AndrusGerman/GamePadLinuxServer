package utils

import "fyne.io/fyne/v2/data/binding"

var DevicesConnect = binding.NewInt()

func init() {
	DevicesConnect.Set(0)
}

func GetDevicesConnect() int {
	v, _ := DevicesConnect.Get()
	return v
}

func SumDevicesConnect(v int) {
	DevicesConnect.Set(GetDevicesConnect() + v)
}
