package server

import "game_pad_linux_server/pkg/devices"

func NewServer(devices devices.Devices) *ServerManagerDefault {
	var ev = new(ServerManagerDefault)
	ev.devices = devices
	return ev
}
