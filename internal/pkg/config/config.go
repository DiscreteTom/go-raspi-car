package config

const (
	// physical pins of the motors
	PWM_A_PIN  = "12" // BCM: 18
	A_IN_1_PIN = "15" // BCM: 22
	A_IN_2_PIN = "13" // BCM: 27
	PWM_B_PIN  = "16" // BCM: 23
	B_IN_1_PIN = "22" // BCM: 25
	B_IN_2_PIN = "18" // BCM: 24

	// physical pins of PS2 controller
	PS2_DAT_PIN = "29" // BCM: 5
	PS2_CMD_PIN = "31" // BCM: 6
	PS2_SEL_PIN = "33" // BCM: 13
	PS2_CLK_PIN = "35" // BCM: 19

	// PS2 controller clock cycle
	PS2_CLK_CYCLE int = 10 // us

	// car
	CAR_MOVE_SPEED = 128
	CAR_TURN_SPEED = 128
)

// computed
const (
	PS2_HALF_CLK_CYCLE int = PS2_CLK_CYCLE >> 1
	PS2_WAIT_INTERVAL  int = PS2_HALF_CLK_CYCLE * 5
)
