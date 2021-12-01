//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package icarous

import (
	"errors"
)

type ICAROUS_FMS_STATE int

const (
	ICAROUS_FMS_STATE_IDLE     ICAROUS_FMS_STATE = 0
	ICAROUS_FMS_STATE_TAKEOFF  ICAROUS_FMS_STATE = 1
	ICAROUS_FMS_STATE_CLIMB    ICAROUS_FMS_STATE = 2
	ICAROUS_FMS_STATE_CRUISE   ICAROUS_FMS_STATE = 3
	ICAROUS_FMS_STATE_APPROACH ICAROUS_FMS_STATE = 4
	ICAROUS_FMS_STATE_LAND     ICAROUS_FMS_STATE = 5
)

var labels_ICAROUS_FMS_STATE = map[ICAROUS_FMS_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_FMS_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ICAROUS_FMS_STATE = map[string]ICAROUS_FMS_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_FMS_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ICAROUS_FMS_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_FMS_STATE) String() string {
	if l, ok := labels_ICAROUS_FMS_STATE[e]; ok {
		return l
	}
	return "invalid value"
}