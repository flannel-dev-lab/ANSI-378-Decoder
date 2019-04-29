package ansi378decoder

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestViewRecords_FMDRecordSmall(t *testing.T) {
	fmd := []byte{1, 2, 3, 4}

	_, err := ViewRecords(fmd, 1)
	if err == nil {
		log.Fatalf("%v", err)
	}
}

// The testcase below should test the validity of the ViewRecords() function in its
// entirety, not just one attribute (length) at a time.
// See record_headers_test.go for an example of table-based testing.

func TestViewRecords(t *testing.T) {
	fmdBase64 := "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="
	fmd, err := base64.StdEncoding.DecodeString(fmdBase64)
	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = RecordHeaders(fmd)
	if err != nil {
		log.Fatalf("%v", err)
	}

/*	if recordHeaders.RecordLength < 65535 {
		viewData, err := ViewRecords(fmd[26:], recordHeaders.Views)
		fmt.Println("a ", viewData, err)
	} else {
		viewData, err := ViewRecords(fmd[30:], recordHeaders.Views)
		fmt.Println("b ", viewData, err)
	}*/
/*	viewData, err := ViewRecords(fmd[26:], recordHeaders.Views)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println(viewData)*/
	//TODO: build test based on decoded data, not length
}

