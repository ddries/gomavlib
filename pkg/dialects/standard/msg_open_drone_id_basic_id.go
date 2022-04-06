//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// Data for filling the OpenDroneID Basic ID message. This and the below messages are primarily meant for feeding data to/from an OpenDroneID implementation. E.g. https://github.com/opendroneid/opendroneid-core-c. These messages are compatible with the ASTM F3411 Remote ID standard and the ASD-STAN prEN 4709-002 Direct Remote ID standard. Additional information and usage of these messages is documented at https://mavlink.io/en/services/opendroneid.html.
type MessageOpenDroneIdBasicId struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Indicates the format for the uas_id field of this message.
	IdType MAV_ODID_ID_TYPE `mavenum:"uint8"`
	// Indicates the type of UA (Unmanned Aircraft).
	UaType MAV_ODID_UA_TYPE `mavenum:"uint8"`
	// UAS (Unmanned Aircraft System) ID following the format specified by id_type. Shall be filled with nulls in the unused portion of the field.
	UasId [20]uint8
}

// GetID implements the msg.Message interface.
func (*MessageOpenDroneIdBasicId) GetID() uint32 {
	return 12900
}
