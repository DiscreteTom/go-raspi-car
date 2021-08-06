package car

import (
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	"fmt"
	"sync"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// drivers
var (
	pwm_a  *gpio.DirectPinDriver
	pwm_b  *gpio.DirectPinDriver
	a_in_1 *gpio.DirectPinDriver
	a_in_2 *gpio.DirectPinDriver
	b_in_1 *gpio.DirectPinDriver
	b_in_2 *gpio.DirectPinDriver
)

// runtime vars
var (
	speedX int16      = 0 // in range [-255, 255]
	speedY int16      = 0 // in range [-255, 255]
	mutex  sync.Mutex     // thread lock
)

func Build(c gobot.Connection) []gobot.Device {
	// construct drivers
	pwm_a = gpio.NewDirectPinDriver(c, config.PWM_A_PIN)
	pwm_b = gpio.NewDirectPinDriver(c, config.PWM_B_PIN)
	a_in_1 = gpio.NewDirectPinDriver(c, config.A_IN_1_PIN)
	a_in_2 = gpio.NewDirectPinDriver(c, config.A_IN_2_PIN)
	b_in_1 = gpio.NewDirectPinDriver(c, config.B_IN_1_PIN)
	b_in_2 = gpio.NewDirectPinDriver(c, config.B_IN_2_PIN)

	return []gobot.Device{pwm_a, pwm_b, a_in_1, a_in_2, b_in_1, b_in_2}
}

func SetSpeedX(x int16) {
	newX := formatSpeed(x)
	if newX != speedX {
		speedX = newX
		updateSpeed()
	}
}

func SetSpeedY(y int16) {
	newY := formatSpeed(y)
	if newY != speedY {
		speedY = newY
		updateSpeed()
	}
}

// change range: [-32768, 32767] => [-127, 127]
func formatSpeed(s int16) int16 {
	ret := s >> 8 // change range: [-32768, 32767] => [-128, 127]
	// avoid overflow
	if ret == -128 {
		ret = -127
	}
	// remove dithering
	if ret > -20 && ret < 20 {
		ret = 0
	}
	return ret
}

func updateSpeed() {
	// thread safe
	mutex.Lock()
	defer mutex.Unlock()

	leftWheelSpeed := speedY - speedX  // range in [-254, 254]
	rightWheelSpeed := speedY + speedX //  // range in [-254, 254]

	applyToWheel(pwm_a, a_in_1, a_in_2, leftWheelSpeed)
	applyToWheel(pwm_b, b_in_1, b_in_2, rightWheelSpeed)
}

func applyToWheel(pwmPin, in1, in2 *gpio.DirectPinDriver, speed int16) {
	if speed == 0 { // stop
		pwmMustWrite(pwmPin, 0)
		digitalMustWrite(in1, 0)
		digitalMustWrite(in2, 0)
	} else if speed < 0 { // move forward
		pwmMustWrite(pwmPin, byte(-speed))
		digitalMustWrite(in1, 1)
		digitalMustWrite(in2, 0)
	} else { // move backward
		pwmMustWrite(pwmPin, byte(speed))
		digitalMustWrite(in1, 0)
		digitalMustWrite(in2, 1)
	}
}

func digitalMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.DigitalWrite(l); err != nil {
		fmt.Println(err)
	}
}

func pwmMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.PwmWrite(l); err != nil {
		fmt.Println(err)
	}
}
