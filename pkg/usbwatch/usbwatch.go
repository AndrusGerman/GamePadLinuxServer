package usbwatch

import (
	"bufio"
	"fmt"
	"game_pad_linux_server/pkg/utils"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			log.Println("signal: ", sig.String())
			os.Exit(0)
		}
	}()

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
	fmt.Println("Close USB Watch")
	if ctx.cmd.Process != nil {
		ctx.cmd.Process.Kill()
	}
}
