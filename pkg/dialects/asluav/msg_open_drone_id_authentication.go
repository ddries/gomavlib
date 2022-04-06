//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Data for filling the OpenDroneID Authentication message. The Authentication Message defines a field that can provide a means of authenticity for the identity of the UAS (Unmanned Aircraft System). The Authentication message can have two different formats. For data page 0, the fields PageCount, Length and TimeStamp are present and AuthData is only 17 bytes. For data page 1 through 15, PageCount, Length and TimeStamp are not present and the size of AuthData is 23 bytes.
type MessageOpenDroneIdAuthentication struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Indicates the type of authentication.
	AuthenticationType MAV_ODID_AUTH_TYPE `mavenum:"uint8"`
	// Allowed range is 0 - 15.
	DataPage uint8
	// This field is only present for page 0. Allowed range is 0 - 15. See the description of struct ODID_Auth_data at https://github.com/opendroneid/opendroneid-core-c/blob/master/libopendroneid/opendroneid.h.
	LastPageIndex uint8
	// This field is only present for page 0. Total bytes of authentication_data from all data pages. See the description of struct ODID_Auth_data at https://github.com/opendroneid/opendroneid-core-c/blob/master/libopendroneid/opendroneid.h.
	Length uint8
	// This field is only present for page 0. 32 bit Unix Timestamp in seconds since 00:00:00 01/01/2019.
	Timestamp uint32
	// Opaque authentication data. For page 0, the size is only 17 bytes. For other pages, the size is 23 bytes. Shall be filled with nulls in the unused portion of the field.
	AuthenticationData [23]uint8
}

// GetID implements the msg.Message interface.
func (*MessageOpenDroneIdAuthentication) GetID() uint32 {
	return 12902
}
