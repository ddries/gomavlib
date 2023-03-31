//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// List of possible failure type to inject.
type FAILURE_TYPE uint32

const (
	// No failure injected, used to reset a previous failure.
	FAILURE_TYPE_OK FAILURE_TYPE = 0
	// Sets unit off, so completely non-responsive.
	FAILURE_TYPE_OFF FAILURE_TYPE = 1
	// Unit is stuck e.g. keeps reporting the same value.
	FAILURE_TYPE_STUCK FAILURE_TYPE = 2
	// Unit is reporting complete garbage.
	FAILURE_TYPE_GARBAGE FAILURE_TYPE = 3
	// Unit is consistently wrong.
	FAILURE_TYPE_WRONG FAILURE_TYPE = 4
	// Unit is slow, so e.g. reporting at slower than expected rate.
	FAILURE_TYPE_SLOW FAILURE_TYPE = 5
	// Data of unit is delayed in time.
	FAILURE_TYPE_DELAYED FAILURE_TYPE = 6
	// Unit is sometimes working, sometimes not.
	FAILURE_TYPE_INTERMITTENT FAILURE_TYPE = 7
)

var labels_FAILURE_TYPE = map[FAILURE_TYPE]string{
	FAILURE_TYPE_OK:           "FAILURE_TYPE_OK",
	FAILURE_TYPE_OFF:          "FAILURE_TYPE_OFF",
	FAILURE_TYPE_STUCK:        "FAILURE_TYPE_STUCK",
	FAILURE_TYPE_GARBAGE:      "FAILURE_TYPE_GARBAGE",
	FAILURE_TYPE_WRONG:        "FAILURE_TYPE_WRONG",
	FAILURE_TYPE_SLOW:         "FAILURE_TYPE_SLOW",
	FAILURE_TYPE_DELAYED:      "FAILURE_TYPE_DELAYED",
	FAILURE_TYPE_INTERMITTENT: "FAILURE_TYPE_INTERMITTENT",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FAILURE_TYPE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_FAILURE_TYPE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FAILURE_TYPE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask FAILURE_TYPE
	for _, label := range labels {
		found := false
		for value, l := range labels_FAILURE_TYPE {
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
func (e FAILURE_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
