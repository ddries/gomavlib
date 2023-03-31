//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

// A mapping of rover flight modes for custom_mode field of heartbeat.
type ROVER_MODE uint32

const (
	ROVER_MODE_MANUAL       ROVER_MODE = 0
	ROVER_MODE_ACRO         ROVER_MODE = 1
	ROVER_MODE_STEERING     ROVER_MODE = 3
	ROVER_MODE_HOLD         ROVER_MODE = 4
	ROVER_MODE_LOITER       ROVER_MODE = 5
	ROVER_MODE_FOLLOW       ROVER_MODE = 6
	ROVER_MODE_SIMPLE       ROVER_MODE = 7
	ROVER_MODE_AUTO         ROVER_MODE = 10
	ROVER_MODE_RTL          ROVER_MODE = 11
	ROVER_MODE_SMART_RTL    ROVER_MODE = 12
	ROVER_MODE_GUIDED       ROVER_MODE = 15
	ROVER_MODE_INITIALIZING ROVER_MODE = 16
)

var labels_ROVER_MODE = map[ROVER_MODE]string{
	ROVER_MODE_MANUAL:       "ROVER_MODE_MANUAL",
	ROVER_MODE_ACRO:         "ROVER_MODE_ACRO",
	ROVER_MODE_STEERING:     "ROVER_MODE_STEERING",
	ROVER_MODE_HOLD:         "ROVER_MODE_HOLD",
	ROVER_MODE_LOITER:       "ROVER_MODE_LOITER",
	ROVER_MODE_FOLLOW:       "ROVER_MODE_FOLLOW",
	ROVER_MODE_SIMPLE:       "ROVER_MODE_SIMPLE",
	ROVER_MODE_AUTO:         "ROVER_MODE_AUTO",
	ROVER_MODE_RTL:          "ROVER_MODE_RTL",
	ROVER_MODE_SMART_RTL:    "ROVER_MODE_SMART_RTL",
	ROVER_MODE_GUIDED:       "ROVER_MODE_GUIDED",
	ROVER_MODE_INITIALIZING: "ROVER_MODE_INITIALIZING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ROVER_MODE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_ROVER_MODE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ROVER_MODE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask ROVER_MODE
	for _, label := range labels {
		found := false
		for value, l := range labels_ROVER_MODE {
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
func (e ROVER_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
