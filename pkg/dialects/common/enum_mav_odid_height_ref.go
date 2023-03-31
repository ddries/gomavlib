//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

type MAV_ODID_HEIGHT_REF uint32

const (
	// The height field is relative to the take-off location.
	MAV_ODID_HEIGHT_REF_OVER_TAKEOFF MAV_ODID_HEIGHT_REF = 0
	// The height field is relative to ground.
	MAV_ODID_HEIGHT_REF_OVER_GROUND MAV_ODID_HEIGHT_REF = 1
)

var labels_MAV_ODID_HEIGHT_REF = map[MAV_ODID_HEIGHT_REF]string{
	MAV_ODID_HEIGHT_REF_OVER_TAKEOFF: "MAV_ODID_HEIGHT_REF_OVER_TAKEOFF",
	MAV_ODID_HEIGHT_REF_OVER_GROUND:  "MAV_ODID_HEIGHT_REF_OVER_GROUND",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_HEIGHT_REF) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_ODID_HEIGHT_REF {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_HEIGHT_REF) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_ODID_HEIGHT_REF
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_ODID_HEIGHT_REF {
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
func (e MAV_ODID_HEIGHT_REF) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
