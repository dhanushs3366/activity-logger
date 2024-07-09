package activitylogger

const (
	DEV_PATH     = "/dev/input/event%d"
	SYS_PATH     = "/sys/class/input/event"
	MAX_INP_DEVS = 1024 //1024 is the max fds linux can handle
)
