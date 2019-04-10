package ANSI_378_Decoder

import (
	"encoding/binary"
	"errors"
)

// Returns the record headers for the FMD
// Format Identifier -> FMR
// Version Number -> 20
// Record Length -> 2 bytes if len is 65535 or 6 bytes if  len is greater than 65535
// CBEFF Owner	-> Vendor specific Owner codes (https://www.ibia.org/cbeff/iso/product-codes)
// Equipment Compliance -> Specifies that device is compliant with FBI standard IAFIS Image Quality Specification, January 29, 1999 of CJIS-RS-0010
// Equipment ID -> Vendor Specified ID, If 0, then contact vendor
// Image Size (X)-> Size of image in X direction
// Image Size (Y)-> Size of image in Y direction
// Image Resolution (X) -> Image Resolution
// Image Resolution (Y) -> Image Resolution
// Views -> Number of finger Views
// Reserved	-> Reserved for future use
func GetFMDRecordHeaders(fmdByteArray []byte) (recordHeaders map[string]interface{}, err error) {
	if len(fmdByteArray) <= 30 {
		return recordHeaders, errors.New("fmd record size is too small")
	}

	recordHeaders = make(map[string]interface{})

	recordHeaders["Format Identifier"] = string(fmdByteArray[0:4])[0:3] // Last byte is null terminator
	recordHeaders["Version Number"] = string(fmdByteArray[4:8])[0:3]    // Last byte is null terminator

	// It means that the record size is 2 bytes
	if fmdByteArray[8] != 0 {
		recordHeaders["Record Length"] = binary.BigEndian.Uint16(fmdByteArray[8:10])
		recordHeaders["CBEFF Owner"] = binary.BigEndian.Uint16(fmdByteArray[10:12])
		recordHeaders["CBEFF Type"] = binary.BigEndian.Uint16(fmdByteArray[12:14])
		recordHeaders["Equipment Compliance"] = binary.BigEndian.Uint16(fmdByteArray[14:16]) >> 12
		recordHeaders["Equipment ID"] = binary.BigEndian.Uint16(fmdByteArray[14:16]) & 4095
		recordHeaders["Image Size (X)"] = binary.BigEndian.Uint16(fmdByteArray[16:18])
		recordHeaders["Image Size (Y)"] = binary.BigEndian.Uint16(fmdByteArray[18:20])
		recordHeaders["Image Resolution (X)"] = binary.BigEndian.Uint16(fmdByteArray[20:22])
		recordHeaders["Image Resolution (Y)"] = binary.BigEndian.Uint16(fmdByteArray[22:24])
		recordHeaders["Views"] = uint8(fmdByteArray[24])
		recordHeaders["Reserved"] = fmdByteArray[25]

	} else { // It means that the record size is 6 bytes
		recordHeaders["Record Length"] = binary.BigEndian.Uint32(fmdByteArray[10:14])
		recordHeaders["CBEFF Owner"] = binary.BigEndian.Uint16(fmdByteArray[14:16])
		recordHeaders["CBEFF Type"] = binary.BigEndian.Uint16(fmdByteArray[16:18])
		recordHeaders["Equipment Compliance"] = binary.BigEndian.Uint16(fmdByteArray[18:20]) >> 12
		recordHeaders["Equipment ID"] = binary.BigEndian.Uint16(fmdByteArray[18:20]) & 4095
		recordHeaders["Image Size (X)"] = binary.BigEndian.Uint16(fmdByteArray[20:22])
		recordHeaders["Image Size (Y)"] = binary.BigEndian.Uint16(fmdByteArray[22:24])
		recordHeaders["Image Resolution (X)"] = binary.BigEndian.Uint16(fmdByteArray[24:26])
		recordHeaders["Image Resolution (Y)"] = binary.BigEndian.Uint16(fmdByteArray[26:28])
		recordHeaders["Views"] = uint8(fmdByteArray[28])
		recordHeaders["Reserved"] = fmdByteArray[29]

	}

	return recordHeaders, nil
}
