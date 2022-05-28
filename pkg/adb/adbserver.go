package adb

import (
	"fmt"
	"game_pad_linux_server/pkg/utils"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/gommon/color"
)

var reverseADBStart = false

func WaitADBClients() {
	fmt.Println(color.Grey("GamePad-adbwath: is start"))
	fmt.Println(color.Grey("GamePad-adbwath: adb allows connection via usb, usb debugging has to be active on your cell phone"))

	go func() {
		for true {
			// Wait Devices
			time.Sleep(time.Second * 2)
			if utils.GetDevicesConnect() > 0 {
				continue
			}
			// IS ADB Device
			devices, err := verifyDeviceConnects()
			if err != nil {
				log.Println("GamePad-adbwath: Brack devices find adb, ", err)
				break
			}

			// adb
			connectReverseAdb(devices)

		}
	}()
}

func connectReverseAdb(devices []string) {
	if len(devices) == 0 {
		reverseADBStart = false
		return
	}
	if reverseADBStart {
		return
	}

	fmt.Println(color.Grey("GamePad-adbwath: List devices Found: "), devices)

	cmd := exec.Command("adb", "reverse", "tcp:8992", "tcp:8992")

	_, err := cmd.Output()
	if err != nil {
		log.Println("GamePad-adbwath: Error create reverse adb ", err)
		log.Println("GamePad-adbwath: If adb is not available the usb connection will not be usable")
		return
	}
	reverseADBStart = true
	fmt.Println(color.Green("GamePad-adbwath: reverse connection complete"))
}

func verifyDeviceConnects() ([]string, error) {

	cmd := exec.Command("adb", "devices", "-l")

	bt, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	spl := strings.Split(string(bt), " ")
	var devices = make([]string, 0)

	// Find Devices
	for _, v := range spl {
		if strings.Contains(v, "model:") {
			devices = append(devices, strings.Split(v, ":")[1])
		}
	}
	// Found Devices
	return devices, nil
}
