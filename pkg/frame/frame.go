// Package frame contains frames and utilities to encode and
// decode them.
package frame

import (
	"bufio"

	"github.com/aler9/gomavlib/pkg/msg"
)

// Frame is the interface implemented by frames of every supported version.
type Frame interface {
	// the system id of the author of the frame.
	GetSystemId() byte

	// the component id of the author of the frame.
	GetComponentId() byte

	// the message encapsuled in the frame.
	GetMessage() msg.Message

	// the frame checksum.
	GetChecksum() uint16

	// generate a clone of the frame
	Clone() Frame

	// decode the frame
	Decode(*bufio.Reader) error

	// encode the frame
	Encode([]byte, []byte) ([]byte, error)

	// generate the checksum
	GenChecksum(byte) uint16
}
