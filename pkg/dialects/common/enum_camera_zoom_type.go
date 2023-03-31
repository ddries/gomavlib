//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Zoom types for MAV_CMD_SET_CAMERA_ZOOM
type CAMERA_ZOOM_TYPE uint32

const (
	// Zoom one step increment (-1 for wide, 1 for tele)
	ZOOM_TYPE_STEP CAMERA_ZOOM_TYPE = 0
	// Continuous zoom up/down until stopped (-1 for wide, 1 for tele, 0 to stop zooming)
	ZOOM_TYPE_CONTINUOUS CAMERA_ZOOM_TYPE = 1
	// Zoom value as proportion of full camera range (a percentage value between 0.0 and 100.0)
	ZOOM_TYPE_RANGE CAMERA_ZOOM_TYPE = 2
	// Zoom value/variable focal length in millimetres. Note that there is no message to get the valid zoom range of the camera, so this can type can only be used for cameras where the zoom range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	ZOOM_TYPE_FOCAL_LENGTH CAMERA_ZOOM_TYPE = 3
)

var labels_CAMERA_ZOOM_TYPE = map[CAMERA_ZOOM_TYPE]string{
	ZOOM_TYPE_STEP:         "ZOOM_TYPE_STEP",
	ZOOM_TYPE_CONTINUOUS:   "ZOOM_TYPE_CONTINUOUS",
	ZOOM_TYPE_RANGE:        "ZOOM_TYPE_RANGE",
	ZOOM_TYPE_FOCAL_LENGTH: "ZOOM_TYPE_FOCAL_LENGTH",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_ZOOM_TYPE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_CAMERA_ZOOM_TYPE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_ZOOM_TYPE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask CAMERA_ZOOM_TYPE
	for _, label := range labels {
		found := false
		for value, l := range labels_CAMERA_ZOOM_TYPE {
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
func (e CAMERA_ZOOM_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
