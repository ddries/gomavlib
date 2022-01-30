//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Settings of a camera. Can be requested with a MAV_CMD_REQUEST_MESSAGE command.
type MessageCameraSettings struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Camera mode
	ModeId CAMERA_MODE `mavenum:"uint8"`
	// Current zoom level (0.0 to 100.0, NaN if not known)
	Zoomlevel float32 `mavext:"true" mavname:"zoomLevel"`
	// Current focus level (0.0 to 100.0, NaN if not known)
	Focuslevel float32 `mavext:"true" mavname:"focusLevel"`
}

// GetID implements the msg.Message interface.
func (*MessageCameraSettings) GetID() uint32 {
	return 260
}