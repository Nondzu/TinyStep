package main

import (
	tinystep "TinyStep"
	"machine"
	"time"
)

func main() {
	leeed()

	estep := tinystep.NewStepEngine()

	stepConfig := tinystep.Config{
		EnablePin: machine.GP18,
		StepPin:   machine.GP17,
		DirPin:    machine.GP16,
		Direction: tinystep.Clockwise,
		Speed:     100,
	}

	estep.Configure(stepConfig)

	for {
		estep.Start(50, 50, tinystep.Clockwise)
		time.Sleep(time.Millisecond * 1200)

		estep.Start(200, 200, tinystep.Anticlockwise)
		time.Sleep(time.Millisecond * 1200)

		estep.Start(150, 750, tinystep.Clockwise)
		time.Sleep(time.Millisecond * 1200)

	}

	// estep.

	// // led := machine.LED
	// ledB := machine.GP22 // blue led
	// ledB.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// enable := machine.GP18 // enable
	// enable.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// step := machine.GP17 // step
	// step.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// step.Low()

	// direction := machine.GP16 // direction
	// direction.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// enable.Low() //!
	// direction.High()

	// dir := false
	// for {

	// 	dir = !dir

	// 	if dir {
	// 		direction.High()
	// 	} else {
	// 		direction.Low()
	// 	}
	// 	// a := 1500
	// 	for i := 0; i < 100; i++ {
	// 		step.High()
	// 		time.Sleep(time.Microsecond * 750)
	// 		step.Low()
	// 		time.Sleep(time.Microsecond * 500)

	// 	}

	// 	ledB.High()

	// 	time.Sleep(time.Millisecond * 50)

	// 	ledB.Low()
	// 	time.Sleep(time.Millisecond * 50)
	// }

}

func leeed() {
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
}
