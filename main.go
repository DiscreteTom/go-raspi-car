package main

import (
	"DiscreteTom/go-raspi-car/internal/pkg/car"
	"fmt"
	"reflect"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "xbox360")

	r := raspi.NewAdaptor()

	carDevices, carFunc := car.Init(r)
	initCar := true

	robot := gobot.NewRobot("test", []gobot.Connection{r, joystickAdaptor}, append(carDevices, stick), func() {
		if initCar {
			carFunc()
			initCar = false
		}

		stick.On(joystick.LeftX, func(data interface{}) {
			// fmt.Println("left_x", data)
			var value, ok = data.(int16)
			if !ok {
				fmt.Println("Data type assert failed: ", reflect.TypeOf(data).Kind())
			}
			car.SetSpeedX(value)

		})
		stick.On(joystick.LeftY, func(data interface{}) {
			// fmt.Println("left_y", data)
			var value, ok = data.(int16)
			if !ok {
				fmt.Println("Data type assert failed: ", reflect.TypeOf(data).Kind())
			}
			car.SetSpeedY(value)
		})

		// for {
		// 	time.Sleep(time.Millisecond * 1000)

		// 	key := ps2.GetKey()
		// 	if key == ps2.NO_KEY {
		// 		car.Stop()
		// 	} else if key == ps2.PAD_UP {
		// 		car.GoForward(config.CAR_MOVE_SPEED)
		// 	} else if key == ps2.PAD_RIGHT {
		// 		car.TurnRight(config.CAR_TURN_SPEED)
		// 	} else if key == ps2.PAD_DOWN {
		// 		car.GoBackward(config.CAR_MOVE_SPEED)
		// 	} else if key == ps2.PAD_LEFT {
		// 		car.TurnLeft(config.CAR_TURN_SPEED)
		// 	}
		// }
	})

	robot.Start()
}
