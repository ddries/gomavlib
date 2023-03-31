//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Flags to report failure cases over the high latency telemtry.
type HL_FAILURE_FLAG uint32

const (
	// GPS failure.
	HL_FAILURE_FLAG_GPS HL_FAILURE_FLAG = 1
	// Differential pressure sensor failure.
	HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE HL_FAILURE_FLAG = 2
	// Absolute pressure sensor failure.
	HL_FAILURE_FLAG_ABSOLUTE_PRESSURE HL_FAILURE_FLAG = 4
	// Accelerometer sensor failure.
	HL_FAILURE_FLAG_3D_ACCEL HL_FAILURE_FLAG = 8
	// Gyroscope sensor failure.
	HL_FAILURE_FLAG_3D_GYRO HL_FAILURE_FLAG = 16
	// Magnetometer sensor failure.
	HL_FAILURE_FLAG_3D_MAG HL_FAILURE_FLAG = 32
	// Terrain subsystem failure.
	HL_FAILURE_FLAG_TERRAIN HL_FAILURE_FLAG = 64
	// Battery failure/critical low battery.
	HL_FAILURE_FLAG_BATTERY HL_FAILURE_FLAG = 128
	// RC receiver failure/no rc connection.
	HL_FAILURE_FLAG_RC_RECEIVER HL_FAILURE_FLAG = 256
	// Offboard link failure.
	HL_FAILURE_FLAG_OFFBOARD_LINK HL_FAILURE_FLAG = 512
	// Engine failure.
	HL_FAILURE_FLAG_ENGINE HL_FAILURE_FLAG = 1024
	// Geofence violation.
	HL_FAILURE_FLAG_GEOFENCE HL_FAILURE_FLAG = 2048
	// Estimator failure, for example measurement rejection or large variances.
	HL_FAILURE_FLAG_ESTIMATOR HL_FAILURE_FLAG = 4096
	// Mission failure.
	HL_FAILURE_FLAG_MISSION HL_FAILURE_FLAG = 8192
)

var labels_HL_FAILURE_FLAG = map[HL_FAILURE_FLAG]string{
	HL_FAILURE_FLAG_GPS:                   "HL_FAILURE_FLAG_GPS",
	HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE: "HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE",
	HL_FAILURE_FLAG_ABSOLUTE_PRESSURE:     "HL_FAILURE_FLAG_ABSOLUTE_PRESSURE",
	HL_FAILURE_FLAG_3D_ACCEL:              "HL_FAILURE_FLAG_3D_ACCEL",
	HL_FAILURE_FLAG_3D_GYRO:               "HL_FAILURE_FLAG_3D_GYRO",
	HL_FAILURE_FLAG_3D_MAG:                "HL_FAILURE_FLAG_3D_MAG",
	HL_FAILURE_FLAG_TERRAIN:               "HL_FAILURE_FLAG_TERRAIN",
	HL_FAILURE_FLAG_BATTERY:               "HL_FAILURE_FLAG_BATTERY",
	HL_FAILURE_FLAG_RC_RECEIVER:           "HL_FAILURE_FLAG_RC_RECEIVER",
	HL_FAILURE_FLAG_OFFBOARD_LINK:         "HL_FAILURE_FLAG_OFFBOARD_LINK",
	HL_FAILURE_FLAG_ENGINE:                "HL_FAILURE_FLAG_ENGINE",
	HL_FAILURE_FLAG_GEOFENCE:              "HL_FAILURE_FLAG_GEOFENCE",
	HL_FAILURE_FLAG_ESTIMATOR:             "HL_FAILURE_FLAG_ESTIMATOR",
	HL_FAILURE_FLAG_MISSION:               "HL_FAILURE_FLAG_MISSION",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HL_FAILURE_FLAG) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_HL_FAILURE_FLAG {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HL_FAILURE_FLAG) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask HL_FAILURE_FLAG
	for _, label := range labels {
		found := false
		for value, l := range labels_HL_FAILURE_FLAG {
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
func (e HL_FAILURE_FLAG) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
