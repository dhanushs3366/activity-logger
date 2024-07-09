package activitylogger

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type keylogger struct {
	File *os.File
}

type DeviceInfo struct {
	DevPath string
	Name    string
}

func New(devPath string) (*keylogger, error) {
	logger := keylogger{}
	inputFile, err := os.OpenFile(devPath, os.O_RDONLY, os.ModeCharDevice)
	if err != nil {
		return nil, err
	}
	logger.File = inputFile
	return &logger, nil
}

func getInputDevPaths() ([]string, error) {

	var devPaths []string
	for i := 0; i < MAX_INP_DEVS; i++ {
		devPath := fmt.Sprintf(DEV_PATH, i)
		if _, err := os.Stat(devPath); errors.Is(err, os.ErrNotExist) {
			continue
		}
		devPaths = append(devPaths, devPath)
	}
	if len(devPaths) == 0 {
		return nil, errors.New("no input devices connected")
	}
	return devPaths, nil
}

func getInputDeviceName(eventNumber int) (string, error) {
	namePath := SYS_PATH + strconv.Itoa(eventNumber) + "/device/name"

	buff, err := os.ReadFile(namePath)
	if err != nil {
		if os.IsPermission(err) {
			return "", errors.New("run with root permissions")
		}
		return "", err
	}

	deviceName := string(buff)
	fmt.Println(deviceName)
	return deviceName, nil
}

func GetInputDevices() ([]DeviceInfo, error) {
	var deviceInfo []DeviceInfo
	devPaths, err := getInputDevPaths()
	if err != nil {
		return nil, err
	}

	for _, devPath := range devPaths {
		eventNo, err := getEventNumber(devPath)
		if err != nil {
			continue
		}
		deviceName, err := getInputDeviceName(eventNo)
		if err != nil {
			continue
		}
		deviceInfo = append(deviceInfo, DeviceInfo{DevPath: devPath, Name: deviceName})
	}
	return deviceInfo, nil
}

func getEventNumber(devPath string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	eventStr := re.FindString(devPath)
	eventNumber, err := strconv.Atoi(eventStr)
	if err != nil {
		return -1, errors.New("provide a proper devPath")
	}
	return eventNumber, nil
}
