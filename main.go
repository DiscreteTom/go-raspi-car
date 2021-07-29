package main

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// init GPIO
	rpio.Open()
	defer rpio.Close()

	initCar()
	initPS2()

	for {
		time.Sleep(time.Millisecond * 60)

		key := getKey()
		if key == NO_KEY {
			stop()
		} else if key == PAD_UP {
			goForward(50)
		} else if key == PAD_RIGHT {
			turnRight(50)
		} else if key == PAD_DOWN {
			goBackward(50)
		} else if key == PAD_LEFT {
			turnLeft(50)
		}
	}
}
