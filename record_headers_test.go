package ANSI_378_Decoder

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestGetFMDRecordHeaders_RecordSizeSmall(t *testing.T) {
	_, err := GetFMDRecordHeaders([]byte{10, 11, 12, 13})

	if err == nil {
		t.Error("err is nil if record size is small")
	}
}

func TestGetFMDRecordHeaders_2ByteRecordLength(t *testing.T) {
	fmdBase64 := "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="

	fmdByteArray, err := base64.StdEncoding.DecodeString(fmdBase64)

	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = GetFMDRecordHeaders(fmdByteArray)
	if err != nil {
		t.Errorf("%v", err)
	}


}

func TestGetFMDRecordHeaders_6ByteRecordLength(t *testing.T) {
	fmdBase64 := "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="

	fmdByteArray, err := base64.StdEncoding.DecodeString(fmdBase64)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmdByteArray[8] = 0

	_, err = GetFMDRecordHeaders(fmdByteArray)
	if err != nil {
		t.Errorf("%v", err)
	}

}