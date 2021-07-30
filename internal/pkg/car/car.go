package car

import (
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

type CarState uint8

const (
	STOPPED CarState = iota
	GO_FORWARD
	GO_BACKWARD
	TURN_LEFT
	TURN_RIGHT
)

var state = STOPPED
var currentSpeed uint32 = 0

var (
	pwm_a rpio.Pin
	pwm_b rpio.Pin

	a_in_1 rpio.Pin
	a_in_2 rpio.Pin
	b_in_1 rpio.Pin
	b_in_2 rpio.Pin
)

func InitCar() {
	pwm_a = rpio.Pin(config.PWM_A_PIN)
	pwm_a.Pwm()
	pwm_a.Freq(100)
	pwm_a.DutyCycle(0, 100)

	pwm_b = rpio.Pin(config.PWM_B_PIN)
	pwm_b.Pwm()
	pwm_b.Freq(100)
	pwm_b.DutyCycle(0, 100)

	a_in_1 = rpio.Pin(config.A_IN_1_PIN)
	a_in_1.Output()
	a_in_2 = rpio.Pin(config.A_IN_2_PIN)
	a_in_2.Output()
	b_in_1 = rpio.Pin(config.B_IN_1_PIN)
	b_in_1.Output()
	b_in_2 = rpio.Pin(config.B_IN_2_PIN)
	b_in_2.Output()
}

func GoForward(speed uint32) {
	if state == GO_FORWARD && currentSpeed == speed {
		return
	}

	state = GO_FORWARD
	currentSpeed = speed
	fmt.Println("move forward with speed ", speed)

	pwm_a.DutyCycle(speed, 100)
	a_in_2.Low()
	a_in_1.High()

	pwm_b.DutyCycle(speed, 100)
	b_in_2.Low()
	b_in_1.High()
}

func Stop() {
	if state == STOPPED {
		return
	}

	state = STOPPED
	fmt.Println("stop")

	pwm_a.DutyCycle(0, 100)
	a_in_1.Low()
	a_in_2.Low()

	pwm_b.DutyCycle(0, 100)
	b_in_1.Low()
	b_in_2.Low()
}

func GoBackward(speed uint32) {
	if state == GO_BACKWARD && currentSpeed == speed {
		return
	}

	state = GO_BACKWARD
	currentSpeed = speed
	fmt.Println("move backward with speed ", speed)

	pwm_a.DutyCycle(speed, 100)
	a_in_2.High()
	a_in_1.Low()

	pwm_b.DutyCycle(speed, 100)
	b_in_2.High()
	b_in_1.Low()
}

func TurnLeft(speed uint32) {
	if state == TURN_LEFT && currentSpeed == speed {
		return
	}

	state = TURN_LEFT
	currentSpeed = speed
	fmt.Println("turn left with speed ", speed)

	pwm_a.DutyCycle(speed, 100)
	a_in_2.High()
	a_in_1.Low()

	pwm_b.DutyCycle(speed, 100)
	b_in_2.Low()
	b_in_1.High()
}

func TurnRight(speed uint32) {
	if state == TURN_RIGHT && currentSpeed == speed {
		return
	}

	state = TURN_RIGHT
	currentSpeed = speed
	fmt.Println("turn right with speed ", speed)

	pwm_a.DutyCycle(speed, 100)
	a_in_2.Low()
	a_in_1.High()

	pwm_b.DutyCycle(speed, 100)
	b_in_2.High()
	b_in_1.Low()
}
