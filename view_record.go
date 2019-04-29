package ansi378

import (
	"encoding/binary"
	"fmt"
)

// ViewRecords returns Minutiae data for each finger view.
// TODO: determine what this means and clean up these docs
// NOTE: I see no conditionals that operate on index length, index 26 or index 30.
// One Finger View Header per each finger view record. For example, 3 views will have 3 Finger View Headers.
// For efficiency, if the record length is < 65535 send fmdByteArray from index 26 to end else send from index 30 to end
// returns a map where key is Finger View {view number} and value is a slice of map where first element of slice is header
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
	//TODO: what is the actual minimum FMD length?
	// if it's less than 5 bytes, then we know there isn't
	// a minutiae record.
	minimumFMDViewRecordLength = 5
)

// ViewRecords returns a slice of ViewRecords given an FMD
func ViewRecords(fmd []byte, views uint8) (*[]ViewRecord, error) {
	if len(fmd) < minimumFMDViewRecordLength {
		return nil, ErrInvalidFMD
	}
	viewRecords := []ViewRecord{}
	var view uint8
	for view = 1; view <= views; view++ {
		viewRecord := ViewRecord{
			FingerPosition: uint8(fmd[0]),
			ViewNumber:     uint8(fmd[1]) >> 4,
			ImpressionType: uint8(fmd[1]) & 15,
			FingerQuality:  uint8(fmd[2]),
			MinutiaeCount:  uint8(fmd[3]),
		}
		fmd = fmd[4:]
		var minutiaes = []Minutiae{}
		var i uint8
		fmt.Println(len(fmd))
		for i = 0; i < viewRecord.MinutiaeCount; i++ {
			record := fmd[6*i : 6*(i+1)]
			m := Minutiae{
				MinutiaeType: binary.BigEndian.Uint16(record[0:2]) >> 14,
				CoordinateX:  binary.BigEndian.Uint16(record[0:2]) & 16383,
				Reserved:     binary.BigEndian.Uint16(record[2:4]) >> 14,
				CoordinateY:  binary.BigEndian.Uint16(record[2:4]) & 16383,
				Angle:        uint8(record[4]),
				Quality:      uint8(record[5]),
			}
			minutiaes = append(minutiaes, m)
		}
		viewRecord.Minutiaes = minutiaes
		viewRecords = append(viewRecords, viewRecord)
	}
	return &viewRecords, nil
}
