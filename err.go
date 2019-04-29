package ansi378decoder

import "errors"

// ErrInvalidFMD occurs when a given FMD byte slice is of invalid length
var ErrInvalidFMD = errors.New("fmd record is invalid")
