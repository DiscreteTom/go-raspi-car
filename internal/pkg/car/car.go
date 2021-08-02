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
	speedX int16      = 0 // [-32768, 32767]
	speedY int16      = 0 // [-32768, 32767]
	mutex  sync.Mutex     // thead lock
)

func Build(c gobot.Connection) ([]gobot.Device, func()) {
	// construct drivers
	pwm_a = gpio.NewDirectPinDriver(c, config.PWM_A_PIN)
	pwm_b = gpio.NewDirectPinDriver(c, config.PWM_B_PIN)
	a_in_1 = gpio.NewDirectPinDriver(c, config.A_IN_1_PIN)
	a_in_2 = gpio.NewDirectPinDriver(c, config.A_IN_2_PIN)
	b_in_1 = gpio.NewDirectPinDriver(c, config.B_IN_1_PIN)
	b_in_2 = gpio.NewDirectPinDriver(c, config.B_IN_2_PIN)

	carInit := func() {
		// init all pin, set default value
		pwm_a.PwmWrite(0)
		pwm_b.PwmWrite(0)
		a_in_1.DigitalWrite(0)
		a_in_2.DigitalWrite(0)
		b_in_1.DigitalWrite(0)
		b_in_2.DigitalWrite(0)
	}

	return []gobot.Device{pwm_a, pwm_b, a_in_1, a_in_2, b_in_1, b_in_2}, carInit
}

func SetSpeedX(x int16) {
	speedX = x
	updateSpeed()
}

func SetSpeedY(y int16) {
	speedY = y
	updateSpeed()
}

func updateSpeed() {
	// thread safe
	mutex.Lock()
	defer mutex.Unlock()

	x := speedX >> 7 // [-32768, 32767] => [-256, 255]
	y := speedY >> 7
	// ensure x, y in [-255, 255]
	if y == -256 {
		y = -255
	}
	if x == -256 {
		x = -255
	}

	if y < 0 { // move forward
		PwmMustWrite(pwm_a, byte(-y))
		DigitalMustWrite(a_in_1, 1)
		DigitalMustWrite(a_in_2, 0)

		PwmMustWrite(pwm_b, byte(-y))
		DigitalMustWrite(b_in_1, 1)
		DigitalMustWrite(b_in_2, 0)
	} else { // move backward
		PwmMustWrite(pwm_a, byte(y))
		DigitalMustWrite(a_in_1, 0)
		DigitalMustWrite(a_in_2, 1)

		PwmMustWrite(pwm_b, byte(y))
		DigitalMustWrite(b_in_1, 0)
		DigitalMustWrite(b_in_2, 1)
	}
}

func DigitalMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.DigitalWrite(l); err != nil {
		fmt.Println(err)
	}
}

func PwmMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.PwmWrite(l); err != nil {
		fmt.Println(err)
	}
}
