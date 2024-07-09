package activitylogger_test

import (
	"fmt"
	"testing"

	activitylogger "github.com/dhanushs3366/activity-logger"
)

// func TestCreate(t *testing.T) {
// 	logger, err := activitylogger.Create("/path/to/device")
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
// 	if logger == nil {
// 		t.Fatalf("Expected logger to be non-nil")
// 	}
// }

func TestGetInputDevices(t *testing.T) {
	deviceInfos, err := activitylogger.GetInputDevices()
	if err != nil {
		t.Fatalf("Expected no err, got %v", err)
	}

	for _, deviceInfo := range deviceInfos {
		fmt.Printf("devPath:%s\t device name:%s", deviceInfo.DevPath, deviceInfo.Name)
	}
}
