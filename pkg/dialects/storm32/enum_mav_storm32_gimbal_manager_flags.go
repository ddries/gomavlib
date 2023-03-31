//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"fmt"
	"strings"
)

// Flags for gimbal manager operation. Used for setting and reporting, unless specified otherwise. If a setting has been accepted by the gimbal manager is reported in the STORM32_GIMBAL_MANAGER_STATUS message.
type MAV_STORM32_GIMBAL_MANAGER_FLAGS uint32

const (
	// 0 = ignore.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_NONE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 0
	// Request to set RC input to active, or report RC input is active. Implies RC mixed. RC exclusive is achieved by setting all clients to inactive.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_RC_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 1
	// Request to set onboard/companion computer client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_ONBOARD_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 2
	// Request to set autopliot client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_AUTOPILOT_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 4
	// Request to set GCS client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 8
	// Request to set camera client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 16
	// Request to set GCS2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 32
	// Request to set camera2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 64
	// Request to set custom client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 128
	// Request to set custom2 client to active, or report this client is active.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM2_ACTIVE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 256
	// Request supervision. This flag is only for setting, it is not reported.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_SUPERVISON MAV_STORM32_GIMBAL_MANAGER_FLAGS = 512
	// Release supervision. This flag is only for setting, it is not reported.
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_RELEASE MAV_STORM32_GIMBAL_MANAGER_FLAGS = 1024
)

var labels_MAV_STORM32_GIMBAL_MANAGER_FLAGS = map[MAV_STORM32_GIMBAL_MANAGER_FLAGS]string{
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_NONE:                    "MAV_STORM32_GIMBAL_MANAGER_FLAGS_NONE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_RC_ACTIVE:               "MAV_STORM32_GIMBAL_MANAGER_FLAGS_RC_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_ONBOARD_ACTIVE:   "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_ONBOARD_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_AUTOPILOT_ACTIVE: "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_AUTOPILOT_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS_ACTIVE:       "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA_ACTIVE:    "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS2_ACTIVE:      "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_GCS2_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA2_ACTIVE:   "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CAMERA2_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM_ACTIVE:    "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM2_ACTIVE:   "MAV_STORM32_GIMBAL_MANAGER_FLAGS_CLIENT_CUSTOM2_ACTIVE",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_SUPERVISON:          "MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_SUPERVISON",
	MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_RELEASE:             "MAV_STORM32_GIMBAL_MANAGER_FLAGS_SET_RELEASE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_MANAGER_FLAGS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_STORM32_GIMBAL_MANAGER_FLAGS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_MANAGER_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_STORM32_GIMBAL_MANAGER_FLAGS
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_STORM32_GIMBAL_MANAGER_FLAGS {
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
func (e MAV_STORM32_GIMBAL_MANAGER_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
