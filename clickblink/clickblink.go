package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"time"

	"machine"
)

const (
	led     = machine.LED
	buttonA = machine.BUTTONA
	buttonB = machine.BUTTONB
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if buttonA.Get() || buttonB.Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
