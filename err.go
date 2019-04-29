package ansi378

import "errors"

// ErrInvalidFMD occurs when a given FMD byte slice is of invalid length
var ErrInvalidFMD = errors.New("fmd record is invalid")
