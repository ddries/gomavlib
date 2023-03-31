//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Possible actions an aircraft can take to avoid a collision.
type MAV_COLLISION_ACTION uint32

const (
	// Ignore any potential collisions
	MAV_COLLISION_ACTION_NONE MAV_COLLISION_ACTION = 0
	// Report potential collision
	MAV_COLLISION_ACTION_REPORT MAV_COLLISION_ACTION = 1
	// Ascend or Descend to avoid threat
	MAV_COLLISION_ACTION_ASCEND_OR_DESCEND MAV_COLLISION_ACTION = 2
	// Move horizontally to avoid threat
	MAV_COLLISION_ACTION_MOVE_HORIZONTALLY MAV_COLLISION_ACTION = 3
	// Aircraft to move perpendicular to the collision's velocity vector
	MAV_COLLISION_ACTION_MOVE_PERPENDICULAR MAV_COLLISION_ACTION = 4
	// Aircraft to fly directly back to its launch point
	MAV_COLLISION_ACTION_RTL MAV_COLLISION_ACTION = 5
	// Aircraft to stop in place
	MAV_COLLISION_ACTION_HOVER MAV_COLLISION_ACTION = 6
)

var labels_MAV_COLLISION_ACTION = map[MAV_COLLISION_ACTION]string{
	MAV_COLLISION_ACTION_NONE:               "MAV_COLLISION_ACTION_NONE",
	MAV_COLLISION_ACTION_REPORT:             "MAV_COLLISION_ACTION_REPORT",
	MAV_COLLISION_ACTION_ASCEND_OR_DESCEND:  "MAV_COLLISION_ACTION_ASCEND_OR_DESCEND",
	MAV_COLLISION_ACTION_MOVE_HORIZONTALLY:  "MAV_COLLISION_ACTION_MOVE_HORIZONTALLY",
	MAV_COLLISION_ACTION_MOVE_PERPENDICULAR: "MAV_COLLISION_ACTION_MOVE_PERPENDICULAR",
	MAV_COLLISION_ACTION_RTL:                "MAV_COLLISION_ACTION_RTL",
	MAV_COLLISION_ACTION_HOVER:              "MAV_COLLISION_ACTION_HOVER",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COLLISION_ACTION) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_COLLISION_ACTION {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COLLISION_ACTION) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_COLLISION_ACTION
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_COLLISION_ACTION {
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
func (e MAV_COLLISION_ACTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
