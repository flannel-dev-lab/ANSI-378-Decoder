package main

import (
	"encoding/base64"
	"fmt"
	ansi378decoder "github.com/flannel-dev-lab/ANSI-378-Decoder"
)

func main() {
	valid := "Rk1SACAyMAABLgAz/v8AAAFlAYgAxQDFAQAAAFYtQFABI3hkgH8BWHZhgFUBQnZgQEoAx3FggJIBLRpfgPwA3k1cQHYBPR1cQEsBTRxcQGAA2BdagO0A2KxaQLABDgpaQG8AkG5ZQK8AuGpZgQUBLZ5ZgMwBBmVYgPgBWJhYQM0AmwRXgMEASWVXQNYAcgVVgHQBE3lUgNEBLaBUQPkAQmBTQDYAlWlTgMABIQRTgJcBRhdTQHoAohBSgKMATwtQgK0BVwVOQKQAhQ9OgJ0BVBdNgFMAgRRLgSYAl01KgDcAdGxGQMYA9QJEQMsBYaBBgB0AoA8+gSYAj6c9QN0AHAg6gBoA+HQ6QBoAsmg3gRQAN7M2ALQBaggxAOcBbTcxAOUBXUErAMsBbU0qAAA="

	//valid := "Rk1SACAyMAAA8gAz/v8AAAFlAYgAxQDFAQAAAFYjgI4A74FigGIANxNfQKoAYgFdQQcBMoldgGkBUSBdQP8BFYlcQJ8AgaxbgO4BC4VZgHQAiRRZgRoAkJRYQLcAdqdXgDwAeBtWgHwAxxdVQLkAYlZVgQcBU35SgKsAPK9SQGgAzR1RgIgBbh9OQN4AkZhNQPsAyZBNQMABD4FNQHQAGHBMgGIAtx1MgLoAiVRLQJAA5HpKgMoAq5NJQJkAuKJIQH0AwxdDQQwBXINDAIoBch00ARoBR40yASkAyZcvAL8AHF4uASYA6jouAJkA4YcqAAA="
	valid2ByteFMD, _    := base64.StdEncoding.DecodeString(valid)

	recordHeader, _ := ansi378decoder.RecordHeaders(valid2ByteFMD)
	fmt.Println("Record length ", recordHeader.RecordLength)
	fmt.Println(ansi378decoder.ViewRecords(valid2ByteFMD[26:], uint8(recordHeader.Views)))
}
