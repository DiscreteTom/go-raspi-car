package car

import (
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	"fmt"
	"sync"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

type Car struct {
	// drivers
	pwmA *gpio.DirectPinDriver
	pwmB *gpio.DirectPinDriver
	aIn1 *gpio.DirectPinDriver
	aIn2 *gpio.DirectPinDriver
	bIn1 *gpio.DirectPinDriver
	bIn2 *gpio.DirectPinDriver

	// runtime vars
	speedX int16      // in range [-255, 255]
	speedY int16      // in range [-255, 255]
	mutex  sync.Mutex // thread lock
}

func NewCar(c gobot.Connection) *Car {
	return &Car{
		pwmA: gpio.NewDirectPinDriver(c, config.PWM_A_PIN),
		pwmB: gpio.NewDirectPinDriver(c, config.PWM_B_PIN),
		aIn1: gpio.NewDirectPinDriver(c, config.A_IN_1_PIN),
		aIn2: gpio.NewDirectPinDriver(c, config.A_IN_2_PIN),
		bIn1: gpio.NewDirectPinDriver(c, config.B_IN_1_PIN),
		bIn2: gpio.NewDirectPinDriver(c, config.B_IN_2_PIN),

		speedX: 0,
		speedY: 0,
	}
}

func (c *Car) Drivers() []gobot.Device {
	return []gobot.Device{c.pwmA, c.pwmB, c.aIn1, c.aIn2, c.bIn1, c.bIn2}
}

func (c *Car) SetSpeedX(x int16) {
	newX := c.formatSpeed(x)
	if newX != c.speedX {
		c.speedX = newX
		c.updateSpeed()
	}
}

func (c *Car) SetSpeedY(y int16) {
	newY := c.formatSpeed(y)
	if newY != c.speedY {
		c.speedY = newY
		c.updateSpeed()
	}
}

// change range: [-32768, 32767] => [-127, 127]
func (c *Car) formatSpeed(s int16) int16 {
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

func (c *Car) updateSpeed() {
	// thread safe
	c.mutex.Lock()
	defer c.mutex.Unlock()

	leftWheelSpeed := c.speedY - c.speedX  // range in [-254, 254]
	rightWheelSpeed := c.speedY + c.speedX //  // range in [-254, 254]

	c.applyToWheel(c.pwmA, c.aIn1, c.aIn2, leftWheelSpeed)
	c.applyToWheel(c.pwmB, c.bIn1, c.bIn2, rightWheelSpeed)
}

func (c *Car) applyToWheel(pwmPin, in1, in2 *gpio.DirectPinDriver, speed int16) {
	if speed == 0 { // stop
		c.pwmMustWrite(pwmPin, 0)
		c.digitalMustWrite(in1, 0)
		c.digitalMustWrite(in2, 0)
	} else if speed < 0 { // move forward
		c.pwmMustWrite(pwmPin, byte(-speed))
		c.digitalMustWrite(in1, 1)
		c.digitalMustWrite(in2, 0)
	} else { // move backward
		c.pwmMustWrite(pwmPin, byte(speed))
		c.digitalMustWrite(in1, 0)
		c.digitalMustWrite(in2, 1)
	}
}

func (c *Car) digitalMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.DigitalWrite(l); err != nil {
		fmt.Println(err)
	}
}

func (c *Car) pwmMustWrite(d *gpio.DirectPinDriver, l byte) {
	if err := d.PwmWrite(l); err != nil {
		fmt.Println(err)
	}
}
