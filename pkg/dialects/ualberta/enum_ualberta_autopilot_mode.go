//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"fmt"
	"strings"
)

// Available autopilot modes for ualberta uav
type UALBERTA_AUTOPILOT_MODE uint32

const (
	// Raw input pulse widts sent to output
	MODE_MANUAL_DIRECT UALBERTA_AUTOPILOT_MODE = 1
	// Inputs are normalized using calibration, the converted back to raw pulse widths for output
	MODE_MANUAL_SCALED UALBERTA_AUTOPILOT_MODE = 2
	MODE_AUTO_PID_ATT  UALBERTA_AUTOPILOT_MODE = 3
	MODE_AUTO_PID_VEL  UALBERTA_AUTOPILOT_MODE = 4
	MODE_AUTO_PID_POS  UALBERTA_AUTOPILOT_MODE = 5
)

var labels_UALBERTA_AUTOPILOT_MODE = map[UALBERTA_AUTOPILOT_MODE]string{
	MODE_MANUAL_DIRECT: "MODE_MANUAL_DIRECT",
	MODE_MANUAL_SCALED: "MODE_MANUAL_SCALED",
	MODE_AUTO_PID_ATT:  "MODE_AUTO_PID_ATT",
	MODE_AUTO_PID_VEL:  "MODE_AUTO_PID_VEL",
	MODE_AUTO_PID_POS:  "MODE_AUTO_PID_POS",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UALBERTA_AUTOPILOT_MODE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_UALBERTA_AUTOPILOT_MODE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UALBERTA_AUTOPILOT_MODE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UALBERTA_AUTOPILOT_MODE
	for _, label := range labels {
		found := false
		for value, l := range labels_UALBERTA_AUTOPILOT_MODE {
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
func (e UALBERTA_AUTOPILOT_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
