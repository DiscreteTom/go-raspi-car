package car

import (
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
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
var currentSpeed byte = 0

var (
	pwm_a *gpio.DirectPinDriver
	pwm_b *gpio.DirectPinDriver

	a_in_1 *gpio.DirectPinDriver
	a_in_2 *gpio.DirectPinDriver
	b_in_1 *gpio.DirectPinDriver
	b_in_2 *gpio.DirectPinDriver
)

func Init(c gobot.Connection) ([]gobot.Device, func()) {
	pwm_a = gpio.NewDirectPinDriver(c, config.PWM_A_PIN)
	pwm_b = gpio.NewDirectPinDriver(c, config.PWM_B_PIN)
	a_in_1 = gpio.NewDirectPinDriver(c, config.A_IN_1_PIN)
	a_in_2 = gpio.NewDirectPinDriver(c, config.A_IN_2_PIN)
	b_in_1 = gpio.NewDirectPinDriver(c, config.B_IN_1_PIN)
	b_in_2 = gpio.NewDirectPinDriver(c, config.B_IN_2_PIN)

	return []gobot.Device{pwm_a, pwm_b, a_in_1, a_in_2, b_in_1, b_in_2}, func() {
		pwm_a.PwmWrite(0)
		pwm_b.PwmWrite(0)
		a_in_1.DigitalWrite(0)
		a_in_2.DigitalWrite(0)
		b_in_1.DigitalWrite(0)
		b_in_2.DigitalWrite(0)
	}
}

func GoForward(speed byte) {
	if state == GO_FORWARD && currentSpeed == speed {
		return
	}

	state = GO_FORWARD
	currentSpeed = speed
	fmt.Println("move forward with speed ", speed)

	pwm_a.PwmWrite(speed)
	a_in_1.DigitalWrite(1)
	a_in_2.DigitalWrite(0)

	pwm_b.PwmWrite(speed)
	b_in_1.DigitalWrite(1)
	b_in_2.DigitalWrite(0)
}

func Stop() {
	if state == STOPPED {
		return
	}

	state = STOPPED
	fmt.Println("stop")

	pwm_a.PwmWrite(0)
	a_in_1.DigitalWrite(0)
	a_in_2.DigitalWrite(0)

	pwm_b.PwmWrite(0)
	b_in_1.DigitalWrite(0)
	b_in_2.DigitalWrite(0)
}

func GoBackward(speed byte) {
	if state == GO_BACKWARD && currentSpeed == speed {
		return
	}

	state = GO_BACKWARD
	currentSpeed = speed
	fmt.Println("move backward with speed ", speed)

	pwm_a.PwmWrite(speed)
	a_in_1.DigitalWrite(0)
	a_in_2.DigitalWrite(1)

	pwm_b.PwmWrite(speed)
	b_in_1.DigitalWrite(0)
	b_in_2.DigitalWrite(1)
}

func TurnLeft(speed byte) {
	if state == TURN_LEFT && currentSpeed == speed {
		return
	}

	state = TURN_LEFT
	currentSpeed = speed
	fmt.Println("turn left with speed ", speed)

	pwm_a.PwmWrite(speed)
	a_in_1.DigitalWrite(0)
	a_in_2.DigitalWrite(1)

	pwm_b.PwmWrite(speed)
	b_in_1.DigitalWrite(1)
	b_in_2.DigitalWrite(0)
}

func TurnRight(speed byte) {
	if state == TURN_RIGHT && currentSpeed == speed {
		return
	}

	state = TURN_RIGHT
	currentSpeed = speed
	fmt.Println("turn right with speed ", speed)

	pwm_a.PwmWrite(speed)
	a_in_1.DigitalWrite(1)
	a_in_2.DigitalWrite(0)

	pwm_b.PwmWrite(speed)
	b_in_1.DigitalWrite(0)
	b_in_2.DigitalWrite(1)
}
