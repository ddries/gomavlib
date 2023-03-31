//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Gimbal device (low level) error flags (bitmap, 0 means no error)
type GIMBAL_DEVICE_ERROR_FLAGS uint32

const (
	// Gimbal device is limited by hardware roll limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 1
	// Gimbal device is limited by hardware pitch limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 2
	// Gimbal device is limited by hardware yaw limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 4
	// There is an error with the gimbal encoders.
	GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 8
	// There is an error with the gimbal power source.
	GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 16
	// There is an error with the gimbal motors.
	GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 32
	// There is an error with the gimbal's software.
	GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 64
	// There is an error with the gimbal's communication.
	GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 128
	// Gimbal device is currently calibrating.
	GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING GIMBAL_DEVICE_ERROR_FLAGS = 256
	// Gimbal device is not assigned to a gimbal manager.
	GIMBAL_DEVICE_ERROR_FLAGS_NO_MANAGER GIMBAL_DEVICE_ERROR_FLAGS = 512
)

var labels_GIMBAL_DEVICE_ERROR_FLAGS = map[GIMBAL_DEVICE_ERROR_FLAGS]string{
	GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT:       "GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT",
	GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT:      "GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT",
	GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT:        "GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT",
	GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR:       "GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR",
	GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR",
	GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR",
	GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR:      "GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR",
	GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR:         "GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR",
	GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING: "GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING",
	GIMBAL_DEVICE_ERROR_FLAGS_NO_MANAGER:          "GIMBAL_DEVICE_ERROR_FLAGS_NO_MANAGER",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_DEVICE_ERROR_FLAGS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GIMBAL_DEVICE_ERROR_FLAGS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_DEVICE_ERROR_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GIMBAL_DEVICE_ERROR_FLAGS
	for _, label := range labels {
		found := false
		for value, l := range labels_GIMBAL_DEVICE_ERROR_FLAGS {
			if l == label {
				mask |= value
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_DEVICE_ERROR_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
