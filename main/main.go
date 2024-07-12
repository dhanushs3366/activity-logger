package main

import (
	"fmt"
	"os"

	activitylogger "github.com/dhanushs3366/activity-logger"
)

func main() {
	devices, err := activitylogger.GetInputDevices()
	if err != nil {
		panic(err)
	}
	var k *activitylogger.Keylogger
	for _, device := range devices {
		fmt.Printf("path: %s\t name:%s", device.DevPath, device.Name)
		if device.DevPath == "/dev/input/event7" {
			k, err = activitylogger.New(device.DevPath)
			if err != nil {
				panic(err)
			}
			if k == nil {
				fmt.Println("No devpath present")
				os.Exit(-1)
			}
			break
		}
	}
	defer k.Close()
	keyInputs := k.Read()

	for input := range keyInputs {
		if input.Value == activitylogger.KEY_PRESSED && input.Type == activitylogger.EV_KEY {
			fmt.Printf("Key type %d\n", input.Type)
			fmt.Printf("key pressed %s \n", input.ToString())
		}

	}
}

/*
asasdfghhjkl;ASXDFAAAAAAAAAAsdffffffffffffg
*/
