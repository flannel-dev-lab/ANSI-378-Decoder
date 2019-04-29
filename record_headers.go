package ansi378

import (
	"encoding/binary"
)

// RecordHeader is header data decoded from an FMD
type RecordHeader struct {
	FormatIdentifier    string
	VersionNumber       string
	RecordLength        uint16
	CBEFFOwner          uint16
	CBEFFType           uint16
	EquipmentCompliance uint16
	EquipmentID         uint16
	ImageSizeX          uint16
	ImageSizeY          uint16
	ImageResolutionX    uint16
	ImageResolutionY    uint16
	Views               uint8
	Reserved            byte
}

const (
	minimumFMDHeaderLength = 30
)

// RecordHeaders returns the record headers of a given FMD
func RecordHeaders(fmd []byte) (*RecordHeader, error) {
	if len(fmd) <= minimumFMDHeaderLength {
		return nil, ErrInvalidFMD
	}
	h := RecordHeader{
		FormatIdentifier: string(fmd[0:4])[0:3], // Last byte is null terminator
		VersionNumber:    string(fmd[4:8])[0:3], // Last byte is null terminator
	}
	// If the record size is 2 bytes
	if fmd[8] != 0 {
		h.RecordLength = binary.BigEndian.Uint16(fmd[8:10])
		h.CBEFFOwner = binary.BigEndian.Uint16(fmd[10:12])
		h.CBEFFType = binary.BigEndian.Uint16(fmd[12:14])
		h.EquipmentCompliance = binary.BigEndian.Uint16(fmd[14:16]) >> 12
		h.EquipmentID = binary.BigEndian.Uint16(fmd[14:16]) & 4095
		h.ImageSizeX = binary.BigEndian.Uint16(fmd[16:18])
		h.ImageSizeY = binary.BigEndian.Uint16(fmd[18:20])
		h.ImageResolutionX = binary.BigEndian.Uint16(fmd[20:22])
		h.ImageResolutionY = binary.BigEndian.Uint16(fmd[22:24])
		h.Views = uint8(fmd[24])
		h.Reserved = fmd[25]

	} else { // The record size is 6 bytes
		h.RecordLength = binary.BigEndian.Uint16(fmd[10:14])
		h.CBEFFOwner = binary.BigEndian.Uint16(fmd[14:16])
		h.CBEFFType = binary.BigEndian.Uint16(fmd[16:18])
		h.EquipmentCompliance = binary.BigEndian.Uint16(fmd[18:20]) >> 12
		h.EquipmentID = binary.BigEndian.Uint16(fmd[18:20]) & 4095
		h.ImageSizeX = binary.BigEndian.Uint16(fmd[20:22])
		h.ImageSizeY = binary.BigEndian.Uint16(fmd[22:26])
		h.ImageResolutionX = binary.BigEndian.Uint16(fmd[24:26])
		h.ImageResolutionY = binary.BigEndian.Uint16(fmd[26:28])
		h.Views = uint8(fmd[28])
		h.Reserved = fmd[29]
	}
	return &h, nil
}
