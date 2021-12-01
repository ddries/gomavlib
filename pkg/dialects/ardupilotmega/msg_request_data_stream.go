//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Request a data stream.
type MessageRequestDataStream struct {
	// The target requested to send the message stream.
	TargetSystem uint8
	// The target requested to send the message stream.
	TargetComponent uint8
	// The ID of the requested data stream
	ReqStreamId uint8
	// The requested message rate
	ReqMessageRate uint16
	// 1 to start sending, 0 to stop sending.
	StartStop uint8
}

// GetID implements the msg.Message interface.
func (*MessageRequestDataStream) GetID() uint32 {
	return 66
}