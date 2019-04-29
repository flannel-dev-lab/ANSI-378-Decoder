package ansi378decoder

import (
	"encoding/base64"
	"testing"
)

var (
	valid2ByteFMDBase64 = "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="
	valid2ByteFMD, _    = base64.StdEncoding.DecodeString(valid2ByteFMDBase64)
	valid2ByteHeader    = RecordHeader{
		CBEFFOwner:          51,
		CBEFFType:           65279,
		EquipmentCompliance: 0,
		EquipmentID:         0,
		FormatIdentifier:    "FMR",
		ImageResolutionX:    197,
		ImageResolutionY:    197,
		ImageSizeX:          357,
		ImageSizeY:          392,
		RecordLength:        302,
		Reserved:            0,
		VersionNumber:       " 20",
		Views:               1,
	}

	valid6ByteFMDBase64 = "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="
	valid6ByteFMD, _    = base64.StdEncoding.DecodeString(valid6ByteFMDBase64)
)

var recordHeaderTests = []struct {
	fmd            []byte
	expectedResult *RecordHeader
	expectedErr    error
}{
	{[]byte{10, 11, 12, 13}, nil, ErrInvalidFMD},
	{valid2ByteFMD, &valid2ByteHeader, nil},
	{valid6ByteFMD, nil, nil},
}

func TestRecordHeaders(t *testing.T) {

	for _, test := range recordHeaderTests {
		res, err := RecordHeaders(test.fmd)
		// if we only expect an error
		if test.expectedResult == nil {
			if err != test.expectedErr {
				t.Errorf("RecordHeaders(%v) => (%v, %v) // Expected (%v, %v)",
					test.fmd, res, err, test.expectedResult, test.expectedErr)
			}
		}
		if res != nil && test.expectedResult != nil {
			if *res != *test.expectedResult || err != test.expectedErr {
				t.Errorf("RecordHeaders(%v) => (%v, %v) // Expected (%v, %v)",
					test.fmd, res, err, test.expectedResult, test.expectedErr)
			}
		}
	}

}
