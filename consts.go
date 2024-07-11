package activitylogger

import (
	"syscall"
)

const (
	DEV_PATH     = "/dev/input/event%d"
	SYS_PATH     = "/sys/class/input/event"
	MAX_INP_DEVS = 1024 //1024 is the max fds linux can handle
)

type InputEvent struct {
	Time  syscall.Timeval //Timestamp of when the event occured
	Type  InputType       //Type of the input event
	Code  uint16          //input event code
	Value int32           // key press
}

const (
	KEY_PRESSED  = 1
	KEY_RELEASED = 0
)

type InputType uint16

const (
	EV_SYN       InputType = 0x00 // Used to sync events
	EV_KEY       InputType = 0x01 // Keyboard and mouse inputs
	EV_REL       InputType = 0x02 // Relative movement
	EV_ABS       InputType = 0x03 // Absolute movement
	EV_MSC       InputType = 0x04 // Miscellaneous events
	EV_SW        InputType = 0x05 // Switch events
	EV_LED       InputType = 0x11 // LED events
	EV_SND       InputType = 0x12 // Sound events
	EV_REP       InputType = 0x14 // Repeating events
	EV_FF        InputType = 0x15 // Force feedback events
	EV_PWR       InputType = 0x16 // Power events
	EV_FF_STATUS InputType = 0x17 // Force feedback status
	EV_MAX       InputType = 0x1f // Maximum value for input events
)
