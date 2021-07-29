package main

type PinNum uint8

const (
	// BCM pins of the motors
	PWM_A_PIN  PinNum = 18
	A_IN_1_PIN PinNum = 22
	A_IN_2_PIN PinNum = 27
	PWM_B_PIN  PinNum = 23
	B_IN_1_PIN PinNum = 25
	B_IN_2_PIN PinNum = 24

	// BCM pins of PS2 controller
	PS2_DAT_PIN PinNum = 5
	PS2_CMD_PIN PinNum = 6
	PS2_SEL_PIN PinNum = 13
	PS2_CLK_PIN PinNum = 19

	// PS2 controller clock cycle
	PS2_CLK_CYCLE uint8 = 20 // us
)

// computed
const (
	PS2_HALF_CLK_CYCLE uint8 = PS2_CLK_CYCLE >> 1
	PS2_WAIT_INTERVAL  uint8 = PS2_HALF_CLK_CYCLE * 5
)
