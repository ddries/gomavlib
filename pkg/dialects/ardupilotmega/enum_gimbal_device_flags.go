//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// Flags for gimbal device (lower level) operation.
type GIMBAL_DEVICE_FLAGS uint32

const (
	// Set to retracted safe position (no stabilization), takes presedence over all other flags.
	GIMBAL_DEVICE_FLAGS_RETRACT GIMBAL_DEVICE_FLAGS = 1
	// Set to neutral/default position, taking precedence over all other flags except RETRACT. Neutral is commonly forward-facing and horizontal (roll=pitch=yaw=0) but may be any orientation.
	GIMBAL_DEVICE_FLAGS_NEUTRAL GIMBAL_DEVICE_FLAGS = 2
	// Lock roll angle to absolute angle relative to horizon (not relative to vehicle). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_ROLL_LOCK GIMBAL_DEVICE_FLAGS = 4
	// Lock pitch angle to absolute angle relative to horizon (not relative to vehicle). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_PITCH_LOCK GIMBAL_DEVICE_FLAGS = 8
	// Lock yaw angle to absolute angle relative to North (not relative to vehicle). If this flag is set, the yaw angle and z component of angular velocity are relative to North (earth frame, x-axis pointing North), else they are relative to the vehicle heading (vehicle frame, earth frame rotated so that the x-axis is pointing forward).
	GIMBAL_DEVICE_FLAGS_YAW_LOCK GIMBAL_DEVICE_FLAGS = 16
	// The gimbal orientation is set exclusively by the RC signals feed to the gimbal's radio control inputs. MAVLink messages for setting the gimbal orientation (GIMBAL_DEVICE_SET_ATTITUDE) are ignored.
	GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE GIMBAL_DEVICE_FLAGS = 256
	// The gimbal orientation is determined by combining/mixing the RC signals feed to the gimbal's radio control inputs and the MAVLink messages for setting the gimbal orientation (GIMBAL_DEVICE_SET_ATTITUDE). How these two controls are combined or mixed is not defined by the protocol but is up to the implementation.
	GIMBAL_DEVICE_FLAGS_RC_MIXED GIMBAL_DEVICE_FLAGS = 512
)

var labels_GIMBAL_DEVICE_FLAGS = map[GIMBAL_DEVICE_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_DEVICE_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_DEVICE_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_DEVICE_FLAGS = map[string]GIMBAL_DEVICE_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_DEVICE_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_DEVICE_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_DEVICE_FLAGS) String() string {
	if l, ok := labels_GIMBAL_DEVICE_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
