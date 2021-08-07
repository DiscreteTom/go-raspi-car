package main

import (
	simpleCar "DiscreteTom/go-raspi-car/internal/pkg/car"
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// platforms
	joystickAdaptor := joystick.NewAdaptor()
	pi := raspi.NewAdaptor()

	// drivers & devices
	stick := joystick.NewDriver(joystickAdaptor, "xbox360")
	car := simpleCar.NewCar(pi)

	robot := gobot.NewRobot("main", []gobot.Connection{pi, joystickAdaptor}, append(car.Drivers(), stick), func() {
		stick.On(joystick.LeftY, func(data interface{}) {
			var value = data.(int16)
			car.SetSpeedY(value)
		})
		stick.On(joystick.RightX, func(data interface{}) {
			var value = data.(int16)
			car.SetSpeedX(value)
		})
	})

	if err := robot.Start(); err != nil { // run and wait for ctrl-c
		fmt.Println(err)
	}
}
