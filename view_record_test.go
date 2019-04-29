package ansi378

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestGetViewRecords_TestRecordLength(t *testing.T) {
	fmdBase64 := "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="

	fmdByteArray, err := base64.StdEncoding.DecodeString(fmdBase64)

	if err != nil {
		log.Fatalf("%v", err)
	}

	recordHeaders, err := GetFMDRecordHeaders(fmdByteArray)

	viewData, err := GetViewRecords(fmdByteArray[26:], recordHeaders["Views"].(uint8))

	if len(viewData["Finger View 1"]) < 46 || len(viewData["Finger View 1"]) > 46 {
		t.Errorf("Got more or less records than expected")
	}
}
