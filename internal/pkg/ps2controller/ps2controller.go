package ps2

import (
	"DiscreteTom/go-raspi-car/internal/pkg/config"
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
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

var MASK = [16]PS2_Key{
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
	ps2_dat *gpio.DirectPinDriver
	ps2_cmd *gpio.DirectPinDriver
	ps2_sel *gpio.DirectPinDriver
	ps2_clk *gpio.DirectPinDriver
)

const (
	HIGH byte = 1
	LOW  byte = 0
)

func Init(c gobot.Connection) ([]gobot.Device, func()) {
	ps2_dat = gpio.NewDirectPinDriver(c, config.PS2_DAT_PIN)

	ps2_cmd = gpio.NewDirectPinDriver(c, config.PS2_CMD_PIN)
	ps2_sel = gpio.NewDirectPinDriver(c, config.PS2_SEL_PIN)
	ps2_clk = gpio.NewDirectPinDriver(c, config.PS2_CLK_PIN)

	work := func() {
		if err := ps2_cmd.DigitalWrite(HIGH); err != nil {
			fmt.Println(err)
		}
		if err := ps2_sel.DigitalWrite(HIGH); err != nil {
			fmt.Println(err)
		}
		if err := ps2_clk.DigitalWrite(HIGH); err != nil {
			fmt.Println(err)
		}
	}

	return []gobot.Device{ps2_cmd, ps2_sel, ps2_clk, ps2_dat}, work
}

func readData(command uint8) uint8 {
	var i uint8
	var j uint8 = 1
	var res uint8 = 0

	for i = 0; i < 8; i++ {
		// send command by bit
		if err := ps2_cmd.DigitalWrite(command & 0x01); err != nil {
			fmt.Println(err)
		}
		command >>= 1
		// and wait for next clock cycle
		time.Sleep(time.Duration(config.PS2_HALF_CLK_CYCLE) * time.Microsecond)

		// set clk to low, send command byte
		if err := ps2_clk.DigitalWrite(LOW); err != nil {
			fmt.Println(err)
		}
		// maintain clock for one cycle
		time.Sleep(time.Duration(config.PS2_HALF_CLK_CYCLE) * time.Microsecond)
		// at the end of this cycle, get input data
		bit, err := ps2_dat.DigitalRead()
		if err != nil {
			fmt.Println(err)
		}
		if bit == 1 {
			res += j
		}
		j <<= 1

		// reset clk to high
		if err := ps2_clk.DigitalWrite(HIGH); err != nil {
			fmt.Println(err)
		}
		// wait for a clock cycle, end this command
		// time.Sleep(time.Duration(config.PS2_HALF_CLK_CYCLE) * time.Microsecond)
	}

	// cool down for a while after sending a byte
	time.Sleep(time.Duration(config.PS2_WAIT_INTERVAL) * time.Microsecond)
	if err := ps2_cmd.DigitalWrite(HIGH); err != nil {
		fmt.Println(err)
	}
	return res
}

func GetKey() PS2_Key {
	var index, i uint8
	var data = [9]uint8{}

	if err := ps2_sel.DigitalWrite(LOW); err != nil {
		fmt.Println(err)
	}
	for i = 0; i < 9; i++ {
		data[i] = readData(scan[i])
	}
	if err := ps2_sel.DigitalWrite(HIGH); err != nil {
		fmt.Println(err)
	}

	var handKey = (uint16(data[4]) << 8) | uint16(data[3])
	for index = 0; index < 16; index++ {
		if (handKey & (1 << (MASK[index] - 1))) == 0 {
			return PS2_Key(index + 1)
		}
	}
	return 0
}
