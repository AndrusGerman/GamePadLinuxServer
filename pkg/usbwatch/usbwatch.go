package usbwatch

import (
	"bufio"
	"fmt"
	"game_pad_linux_server/pkg/utils"
	"io"
	"os/exec"

	"github.com/labstack/gommon/color"
)

type USBWatch struct {
	cmd  *exec.Cmd
	pipe io.ReadCloser
}

func NewUSBWatch() (*USBWatch, error) {
	cmd := exec.Command("udevadm", "monitor", "--subsystem-match=usb", "--udev")
	salida, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil
	}
	err = cmd.Start()
	if err != nil {
		return nil, nil
	}

	return &USBWatch{
		cmd:  cmd,
		pipe: salida,
	}, nil
}

func (ctx *USBWatch) WatchOn(connect func(), disconnect func()) {
	scanner := bufio.NewScanner(ctx.pipe)
	for scanner.Scan() {
		text := scanner.Text()
		if utils.MultipleContains(text, " bind", "(usb)") {
			connect()
		}
		if utils.MultipleContains(text, " unbind", "(usb)") {
			fmt.Println(color.Grey("GamePad-usbwatch: some USB device was disconnected"))
			disconnect()
		}
	}
}

func (ctx *USBWatch) Close() {
	if ctx.cmd.Process != nil {
		ctx.cmd.Process.Kill()
	}
}
