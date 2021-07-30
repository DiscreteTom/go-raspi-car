package main

import (
	"DiscreteTom/go-raspi-car/internal/pkg/car"
	ps2 "DiscreteTom/go-raspi-car/internal/pkg/ps2controller"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// init GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
	}
	defer rpio.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	car.InitCar()
	ps2.InitPS2()

	for {
		time.Sleep(time.Millisecond * 1000)

		var loop = true
		select {
		case <-c:
			fmt.Println("got keyboard interrupt")
			loop = false
		default:
		}
		if !loop {
			break
		}

		key := ps2.GetKey()
		if key == ps2.NO_KEY {
			car.Stop()
		} else if key == ps2.PAD_UP {
			car.GoForward(50)
		} else if key == ps2.PAD_RIGHT {
			car.TurnRight(50)
		} else if key == ps2.PAD_DOWN {
			car.GoBackward(50)
		} else if key == ps2.PAD_LEFT {
			car.TurnLeft(50)
		}
	}
}
