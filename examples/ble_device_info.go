package main

import (
	"fmt"
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
)

func main() {
	gbot := gobot.NewMaster()

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	info := ble.NewDeviceInformationDriver(bleAdaptor)

	work := func() {
		fmt.Println("Model number:", info.GetModelNumber())
		fmt.Println("Firmware rev:", info.GetFirmwareRevision())
		fmt.Println("Hardware rev:", info.GetHardwareRevision())
		fmt.Println("Manufacturer name:", info.GetManufacturerName())
		fmt.Println("PnPId:", info.GetPnPId())
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{info},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
