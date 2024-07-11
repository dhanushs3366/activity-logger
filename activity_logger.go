package activitylogger

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unsafe"
)

type Keylogger struct {
	File *os.File
}

type DeviceInfo struct {
	DevPath string
	Name    string
}

func New(devPath string) (*Keylogger, error) {
	logger := Keylogger{}
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

func (k *Keylogger) Read() chan InputEvent {
	inputChannel := make(chan InputEvent)

	go func() {
		for {
			input, err := k.parseBinaryKeyCode()
			if err != nil {
				close(inputChannel)
				break
			}
			if inputChannel != nil {
				inputChannel <- *input
			}

		}
	}()

	return inputChannel
}

func (k *Keylogger) parseBinaryKeyCode() (*InputEvent, error) {
	eventSize := int(unsafe.Sizeof(InputEvent{}))
	buff := make([]byte, eventSize)
	_, err := k.File.Read(buff)
	if err != nil {
		return nil, err
	}

	var keyInput InputEvent
	err = binary.Read(bytes.NewBuffer(buff), binary.LittleEndian, &keyInput)
	if err != nil {
		return nil, err
	}
	return &keyInput, nil
}

func (i *InputEvent) ToString() string {
	return KeyMap[int(i.Code)]
}

func (k *Keylogger) Close() error {
	return k.File.Close()
}
