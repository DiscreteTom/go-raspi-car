package main

import (
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

	initCar()
	initPS2()

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
