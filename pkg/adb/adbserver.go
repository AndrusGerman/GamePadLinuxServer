package adb

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

var reverseADBStart = false
var DevicesConnect = 0

func WaitADBClients() {
	log.Println("adbwatch: start")
	log.Println("adbwatch: adb allows connection via usb, usb debugging has to be active on your cell phone")

	go func() {
		for true {
			// Wait Devices
			time.Sleep(time.Second * 2)
			if DevicesConnect > 0 {
				continue
			}
			// IS ADB Device
			devices, err := verifyDeviceConnects()
			if err != nil {
				log.Println("adbwatch: Brack devices find adb, ", err)
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

	log.Println("adbwatch: List devices Found: ", devices)

	cmd := exec.Command("adb", "reverse", "tcp:8992", "tcp:8992")

	_, err := cmd.Output()
	if err != nil {
		log.Println("adbwatch: Error create reverse adb ", err)
		log.Println("adbwatch: If adb is not available the usb connection will not be usable")
		return
	}
	reverseADBStart = true
	log.Println("adbwatch: reverse connection complete")
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
