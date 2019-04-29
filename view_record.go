package ansi378decoder

import (
	"encoding/binary"
)

// Returns the Minutiae data for each finger view
// One Finger View Header per each finger view record. For example, 3 views will have 3  Finger View Headers
// For efficiency, if the record length is < 65535 send fmdByteArray from index 26 to end else send from index 30 to end
// Returns a map where key is Finger View {view number} and value is a slice of map where first element of slice is header
// and rest is minutiae data

// ViewRecord is data constructed from a FMD
type ViewRecord struct {
	FingerPosition uint8
	ViewNumber     uint8
	ImpressionType uint8
	FingerQuality  uint8
	MinutiaeCount  uint8
	Minutiaes      []Minutiae
}

// Minutiae samples are included in each ViewRecord
type Minutiae struct {
	MinutiaeType uint16
	CoordinateX  uint16
	CoordinateY  uint16
	Reserved     uint16
	Angle        uint8
	Quality      uint8
}

const (
	minimumFMDViewRecordLength = 5
)

// ViewRecords returns a slice of ViewRecords given an FMD
func ViewRecords(fmd []byte, views uint8) ([]ViewRecord, error) {
	if len(fmd) < minimumFMDViewRecordLength {
		return nil, ErrInvalidFMD
	}

	var view uint8
	var viewRecords []ViewRecord

	for view = 1; view <= views; view++ {

		viewRecord := ViewRecord{
			FingerPosition: uint8(fmd[0]),
			ViewNumber:     uint8(fmd[1]) >> 4,
			ImpressionType: uint8(fmd[1]) & 15,
			FingerQuality:  uint8(fmd[2]),
			MinutiaeCount:  uint8(fmd[3]),
		}

		fmd = fmd[4:]

		var minutiaeCount uint

		var minutiaes []Minutiae

		for minutiaeCount = 0; minutiaeCount < uint(viewRecord.MinutiaeCount); minutiaeCount++ {
			record := fmd[6*minutiaeCount : 6*(minutiaeCount+1)]

			minutiae := Minutiae{
				MinutiaeType: binary.BigEndian.Uint16(record[0:2]) >> 14,
				CoordinateX:  binary.BigEndian.Uint16(record[0:2]) & 16383,
				Reserved:     binary.BigEndian.Uint16(record[2:4]) >> 14,
				CoordinateY:  binary.BigEndian.Uint16(record[2:4]) & 16383,
				Angle:        uint8(record[4]),
				Quality:      uint8(record[5]),
			}
			minutiaes = append(minutiaes, minutiae)
		}
		viewRecord.Minutiaes = minutiaes
		viewRecords = append(viewRecords, viewRecord)
	}
	return viewRecords, nil
}
