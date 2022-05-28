package utils

import "fyne.io/fyne/v2/data/binding"

func init() {
	DevicesConnect.Set(0)
	DevicesList.Set(make([]string, 0))
}

// Devices Connect
var DevicesConnect = binding.NewInt()

func GetDevicesConnect() int {
	v, _ := DevicesConnect.Get()
	return v
}

func SumDevicesConnect(v int) {
	DevicesConnect.Set(GetDevicesConnect() + v)
}

// Device List connect
var DevicesList = binding.NewStringList()

func GetDevicesList() []string {
	v, _ := DevicesList.Get()
	return v
}
