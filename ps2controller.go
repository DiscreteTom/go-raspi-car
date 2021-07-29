package main

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type PS2_Key uint8

const (
	NO_KEY PS2_Key = iota
	SELECT
	L3
	R3
	START
	PAD_UP
	PAD_RIGHT
	PAD_DOWN
	PAD_LEFT
	L2
	R2
	L1
	R1
	TRIANGLE
	CIRCLE
	CROSS
	SQUARE
)

var MASK = [17]PS2_Key{
	NO_KEY,
	SELECT,
	L3,
	R3,
	START,
	PAD_UP,
	PAD_RIGHT,
	PAD_DOWN,
	PAD_LEFT,
	L2,
	R2,
	L1,
	R1,
	TRIANGLE,
	CIRCLE,
	CROSS,
	SQUARE,
}

var scan = [9]uint8{0x01, 0x42, 0, 0, 0, 0, 0, 0, 0}

var (
	ps2_dat rpio.Pin
	ps2_cmd rpio.Pin
	ps2_sel rpio.Pin
	ps2_clk rpio.Pin
)

func initPS2() {
	ps2_dat = rpio.Pin(PS2_DAT_PIN)
	ps2_dat.Input()

	ps2_cmd = rpio.Pin(PS2_CMD_PIN)
	ps2_cmd.Output()
	ps2_cmd.High()
	ps2_sel = rpio.Pin(PS2_SEL_PIN)
	ps2_sel.Output()
	ps2_sel.High()
	ps2_clk = rpio.Pin(PS2_CLK_PIN)
	ps2_clk.Output()
	ps2_clk.High()
}

func readData(command uint8) uint8 {
	var i uint8
	var j uint8 = 1
	var res uint8 = 0

	for i = 0; i < 8; i++ {
		// send command by bit
		if command&0x01 == 1 {
			ps2_cmd.High()
		} else {
			ps2_cmd.Low()
		}
		command >>= 1
		// and wait for next clock cycle
		time.Sleep(time.Duration(PS2_CLK_CYCLE) * time.Microsecond)

		// set clk to low, send command byte
		ps2_clk.Low()
		// maintain clock for one cycle
		time.Sleep(time.Duration(PS2_CLK_CYCLE) * time.Microsecond)
		// at the end of this cycle, get input data
		if ps2_dat.Read() == rpio.High {
			res += j
		}
		j <<= 1

		// reset clk to high
		ps2_clk.High()
		// wait for a clock cycle, end this command
		time.Sleep(time.Duration(PS2_CLK_CYCLE) * time.Microsecond)
	}

	// coll down for a while after sending a byte
	time.Sleep(time.Duration(PS2_WAIT_INTERVAL) * time.Microsecond)
	ps2_cmd.High()
	return res
}

func getKey() PS2_Key {
	var index, i uint8
	var data = [9]uint8{}

	ps2_sel.Low()
	for i = 0; i < 9; i++ {
		data[i] = readData(scan[i])
	}
	ps2_sel.High()

	var handKey = (data[4] << 8) | data[3]
	for index = 0; index < 16; index++ {
		if handKey&(1<<(MASK[index]-1)) == 0 {
			return PS2_Key(index + 1)
		}
	}
	return 0
}
