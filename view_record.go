package ansi378

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Returns the Minutiae data for each finger view
// One Finger View Header per each finger view record. For example, 3 views will have 3  Finger View Headers
// For efficiency, if the record length is < 65535 send fmdByteArray from index 26 to end else send from index 30 to end
// Returns a map where key is Finger View {view number} and value is a slice of map where first element of slice is header
// and rest is minutiae data

func GetViewRecords(fmdByteArray []byte, views uint8) (fingerViewRecords map[string][]interface{}, err error) {
	var view uint8

	if len(fmdByteArray) == 0 {
		return fingerViewRecords, errors.New("empty fmd byte array")
	}

	fingerViewRecords = make(map[string][]interface{})

	for view = 1; view <= views; view++ {
		fingerViewHeader := make(map[string]uint8)

		fingerViewHeader["Finger Position"] = uint8(fmdByteArray[0])
		fingerViewHeader["View Number"] = uint8(fmdByteArray[1]) >> 4
		fingerViewHeader["Impression Type"] = uint8(fmdByteArray[1]) & 15
		fingerViewHeader["Finger Quality"] = uint8(fmdByteArray[2])
		fingerViewHeader["Minutiae Count"] = uint8(fmdByteArray[3])

		fingerViewRecords[fmt.Sprintf("Finger View %d", view)] = append(fingerViewRecords[fmt.Sprintf("Finger View %d", view)], fingerViewHeader)

		fmdByteArray = fmdByteArray[4:]

		var minutiaeCount uint
		for minutiaeCount = 0; minutiaeCount < uint(fingerViewHeader["Minutiae Count"]); minutiaeCount++ {

			minutiaeRecord := fmdByteArray[6*minutiaeCount : 6*(minutiaeCount+1)]

			minutiaeData := make(map[string]interface{})
			minutiaeData["Minutiae Type"] = binary.BigEndian.Uint16(minutiaeRecord[0:2]) >> 14
			minutiaeData["X coordinate"] = binary.BigEndian.Uint16(minutiaeRecord[0:2]) & 16383
			minutiaeData["Reserved"] = binary.BigEndian.Uint16(minutiaeRecord[2:4]) >> 14
			minutiaeData["Y coordinate"] = binary.BigEndian.Uint16(minutiaeRecord[2:4]) & 16383
			minutiaeData["Angle"] = uint8(minutiaeRecord[4])
			minutiaeData["Quality"] = uint8(minutiaeRecord[5])

			fingerViewRecords[fmt.Sprintf("Finger View %d", view)] = append(fingerViewRecords[fmt.Sprintf("Finger View %d", view)], minutiaeData)
		}
	}

	return fingerViewRecords, nil
}
