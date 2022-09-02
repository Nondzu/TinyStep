// Package TinyStep implements a step engine library for the A4988 step engine driver
//
//

/*
todo:
1.  return channel fish finish flag from run command

*/

package TinyStep

import (
	"machine"
	"time"
)

type StepEngine struct {
	enablePin machine.Pin // enable GPIO
	stepPin   machine.Pin // step GPIO
	dirPin    machine.Pin // direction GPIO
	speed     uint16      // speed in step/s default 200step/s
	direction Direction   // default Clockwise
	enable    bool        // default Clockwise
}

// Config is the configuration for the display
type Config struct {
	EnablePin machine.Pin
	StepPin   machine.Pin
	DirPin    machine.Pin
	Direction Direction
	Speed     uint16
}

// Engine Direction
const (
	Clockwise     = 0x01
	Anticlockwise = 0x02
)

type Direction uint8

// NewI2C creates a new SSD1306 connection. The I2C wire must already be configured.
func NewStepEngine() StepEngine {
	return StepEngine{}
}

// Configure initializes the display with default configuration
func (s *StepEngine) Configure(cfg Config) {

	s.enablePin = cfg.EnablePin
	s.stepPin = cfg.StepPin
	s.dirPin = cfg.DirPin

	s.enablePin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.stepPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.dirPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	//set direction
	if cfg.Direction != 0 {
		s.direction = cfg.Direction
	} else {
		s.direction = Clockwise
	}

	s.speed = cfg.Speed
	s.enable = true
}

func (s *StepEngine) SetDirection(dir Direction) {

	if (dir != 0) && (dir <= Anticlockwise) {
		s.direction = dir
	} else {
		s.direction = Clockwise
	}

	if s.direction == Clockwise {
		s.dirPin.High()
	} else if s.direction == Anticlockwise {
		s.dirPin.Low()
	}
}

func (s *StepEngine) Start(steps uint32, speed uint16, dir Direction) {

	s.enableEngine(true)

	//200
	/*
	   calculate speed:
	   200 step/s
	   1s = 1000 ms
	   1ms = 1000 us


	   1 000 000 / 200

	*/
	// s.dirPin.High()

	s.SetDirection(dir)

	s.stepPin.Low()
	stepTime := time.Duration(1000000/uint32(speed)-50) * time.Microsecond // 1 000 000 microsecond / speed (steps/s) / 2

	_ = stepTime
	go func() {
		for i := uint32(0); i < steps; i++ {
			s.stepPin.High()
			time.Sleep(time.Microsecond * 50)
			s.stepPin.Low()
			time.Sleep(stepTime)
		}
	}()

	//return channel
}

func (s *StepEngine) enableEngine(e bool) {

	if e == true {
		s.enablePin.Low()
	} else {
		s.enablePin.High()
	}
}

// func (s *StepEngine) SetDirection(e bool) {

// 	if e == true {
// 		s.enablePin.Low()
// 	} else {
// 		s.enablePin.High()
// 	}
// }
