//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// SERIAL_CONTROL flags (bitmask)
type SERIAL_CONTROL_FLAG uint32

const (
	// Set if this is a reply
	SERIAL_CONTROL_FLAG_REPLY SERIAL_CONTROL_FLAG = 1
	// Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message
	SERIAL_CONTROL_FLAG_RESPOND SERIAL_CONTROL_FLAG = 2
	// Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set
	SERIAL_CONTROL_FLAG_EXCLUSIVE SERIAL_CONTROL_FLAG = 4
	// Block on writes to the serial port
	SERIAL_CONTROL_FLAG_BLOCKING SERIAL_CONTROL_FLAG = 8
	// Send multiple replies until port is drained
	SERIAL_CONTROL_FLAG_MULTI SERIAL_CONTROL_FLAG = 16
)

var labels_SERIAL_CONTROL_FLAG = map[SERIAL_CONTROL_FLAG]string{
	SERIAL_CONTROL_FLAG_REPLY:     "SERIAL_CONTROL_FLAG_REPLY",
	SERIAL_CONTROL_FLAG_RESPOND:   "SERIAL_CONTROL_FLAG_RESPOND",
	SERIAL_CONTROL_FLAG_EXCLUSIVE: "SERIAL_CONTROL_FLAG_EXCLUSIVE",
	SERIAL_CONTROL_FLAG_BLOCKING:  "SERIAL_CONTROL_FLAG_BLOCKING",
	SERIAL_CONTROL_FLAG_MULTI:     "SERIAL_CONTROL_FLAG_MULTI",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SERIAL_CONTROL_FLAG) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_SERIAL_CONTROL_FLAG {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SERIAL_CONTROL_FLAG) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask SERIAL_CONTROL_FLAG
	for _, label := range labels {
		found := false
		for value, l := range labels_SERIAL_CONTROL_FLAG {
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
func (e SERIAL_CONTROL_FLAG) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
