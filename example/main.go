package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			led.High()
			time.Sleep(time.Millisecond * 500)
			led.Low()
		}

	}()

	// led := machine.LED
	ledB := machine.GP22 // blue led
	ledB.Configure(machine.PinConfig{Mode: machine.PinOutput})

	enable := machine.GP18 // enable
	enable.Configure(machine.PinConfig{Mode: machine.PinOutput})

	step := machine.GP17 // step
	step.Configure(machine.PinConfig{Mode: machine.PinOutput})
	step.Low()

	direction := machine.GP16 // direction
	direction.Configure(machine.PinConfig{Mode: machine.PinOutput})

	enable.Low() //!
	direction.High()

	dir := false
	for {

		dir = !dir

		if dir {
			direction.High()
		} else {
			direction.Low()
		}
		// a := 1500
		for i := 0; i < 150; i++ {
			step.High()
			time.Sleep(time.Microsecond * 750)
			step.Low()
			time.Sleep(time.Microsecond * 500)

		}

		ledB.High()

		time.Sleep(time.Millisecond * 200)

		ledB.Low()
		time.Sleep(time.Millisecond * 200)
	}

}
