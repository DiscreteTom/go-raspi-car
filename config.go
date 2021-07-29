package main

type PinNum uint8

const (
	// physical pins of the motors
	PWM_A_PIN  PinNum = 12 // BCM: 18
	A_IN_1_PIN PinNum = 15 // BCM: 22
	A_IN_2_PIN PinNum = 13 // BCM: 27
	PWM_B_PIN  PinNum = 16 // BCM: 23
	B_IN_1_PIN PinNum = 22 // BCM: 25
	B_IN_2_PIN PinNum = 18 // BCM: 24

	// physical pins of PS2 controller
	PS2_DAT_PIN PinNum = 29 // BCM: 5
	PS2_CMD_PIN PinNum = 31 // BCM: 6
	PS2_SEL_PIN PinNum = 33 // BCM: 13
	PS2_CLK_PIN PinNum = 35 // BCM: 19

	// PS2 controller clock cycle
	PS2_CLK_CYCLE uint8 = 20 // us
)

// computed
const (
	PS2_HALF_CLK_CYCLE uint8 = PS2_CLK_CYCLE >> 1
	PS2_WAIT_INTERVAL  uint8 = PS2_HALF_CLK_CYCLE * 5
)
