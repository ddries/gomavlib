//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"fmt"
	"strings"
)

// Enumeration of possible shot modes.
type MAV_QSHOT_MODE uint32

const (
	// Undefined shot mode. Can be used to determine if qshots should be used or not.
	MAV_QSHOT_MODE_UNDEFINED MAV_QSHOT_MODE = 0
	// Start normal gimbal operation. Is usually used to return back from a shot.
	MAV_QSHOT_MODE_DEFAULT MAV_QSHOT_MODE = 1
	// Load and keep safe gimbal position and stop stabilization.
	MAV_QSHOT_MODE_GIMBAL_RETRACT MAV_QSHOT_MODE = 2
	// Load neutral gimbal position and keep it while stabilizing.
	MAV_QSHOT_MODE_GIMBAL_NEUTRAL MAV_QSHOT_MODE = 3
	// Start mission with gimbal control.
	MAV_QSHOT_MODE_GIMBAL_MISSION MAV_QSHOT_MODE = 4
	// Start RC gimbal control.
	MAV_QSHOT_MODE_GIMBAL_RC_CONTROL MAV_QSHOT_MODE = 5
	// Start gimbal tracking the point specified by Lat, Lon, Alt.
	MAV_QSHOT_MODE_POI_TARGETING MAV_QSHOT_MODE = 6
	// Start gimbal tracking the system with specified system ID.
	MAV_QSHOT_MODE_SYSID_TARGETING MAV_QSHOT_MODE = 7
	// Start 2-point cable cam quick shot.
	MAV_QSHOT_MODE_CABLECAM_2POINT MAV_QSHOT_MODE = 8
	// Start gimbal tracking the home location.
	MAV_QSHOT_MODE_HOME_TARGETING MAV_QSHOT_MODE = 9
)

var labels_MAV_QSHOT_MODE = map[MAV_QSHOT_MODE]string{
	MAV_QSHOT_MODE_UNDEFINED:         "MAV_QSHOT_MODE_UNDEFINED",
	MAV_QSHOT_MODE_DEFAULT:           "MAV_QSHOT_MODE_DEFAULT",
	MAV_QSHOT_MODE_GIMBAL_RETRACT:    "MAV_QSHOT_MODE_GIMBAL_RETRACT",
	MAV_QSHOT_MODE_GIMBAL_NEUTRAL:    "MAV_QSHOT_MODE_GIMBAL_NEUTRAL",
	MAV_QSHOT_MODE_GIMBAL_MISSION:    "MAV_QSHOT_MODE_GIMBAL_MISSION",
	MAV_QSHOT_MODE_GIMBAL_RC_CONTROL: "MAV_QSHOT_MODE_GIMBAL_RC_CONTROL",
	MAV_QSHOT_MODE_POI_TARGETING:     "MAV_QSHOT_MODE_POI_TARGETING",
	MAV_QSHOT_MODE_SYSID_TARGETING:   "MAV_QSHOT_MODE_SYSID_TARGETING",
	MAV_QSHOT_MODE_CABLECAM_2POINT:   "MAV_QSHOT_MODE_CABLECAM_2POINT",
	MAV_QSHOT_MODE_HOME_TARGETING:    "MAV_QSHOT_MODE_HOME_TARGETING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_QSHOT_MODE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_QSHOT_MODE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_QSHOT_MODE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_QSHOT_MODE
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_QSHOT_MODE {
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
func (e MAV_QSHOT_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
