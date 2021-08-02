package main

import (
	"DiscreteTom/go-raspi-car/internal/pkg/car"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	joystickAdaptor := joystick.NewAdaptor()
	pi := raspi.NewAdaptor()

	stick := joystick.NewDriver(joystickAdaptor, "xbox360")
	carDevices, carInit := car.Build(pi)

	robot := gobot.NewRobot("test", []gobot.Connection{pi, joystickAdaptor}, append(carDevices, stick), func() {
		carInit()

		stick.On(joystick.LeftX, func(data interface{}) {
			var value = data.(int16)
			car.SetSpeedX(value)
		})
		stick.On(joystick.LeftY, func(data interface{}) {
			var value = data.(int16)
			car.SetSpeedY(value)
		})
	})

	robot.Start()
}
