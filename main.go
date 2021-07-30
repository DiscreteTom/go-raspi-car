package main

import (
	"DiscreteTom/go-raspi-car/internal/pkg/car"
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	ps2 "DiscreteTom/go-raspi-car/internal/pkg/ps2controller"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	carDevices, carFunc := car.Init(r)
	ps2Devices, ps2Func := ps2.Init(r)

	robot := gobot.NewRobot("test", []gobot.Connection{r}, append(ps2Devices, carDevices...), func() {
		carFunc()
		ps2Func()

		for {
			time.Sleep(time.Millisecond * 1000)

			key := ps2.GetKey()
			if key == ps2.NO_KEY {
				car.Stop()
			} else if key == ps2.PAD_UP {
				car.GoForward(config.CAR_MOVE_SPEED)
			} else if key == ps2.PAD_RIGHT {
				car.TurnRight(config.CAR_TURN_SPEED)
			} else if key == ps2.PAD_DOWN {
				car.GoBackward(config.CAR_MOVE_SPEED)
			} else if key == ps2.PAD_LEFT {
				car.TurnLeft(config.CAR_TURN_SPEED)
			}
		}
	})

	robot.Start()
}
