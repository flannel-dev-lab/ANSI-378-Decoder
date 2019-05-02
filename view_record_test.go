package ansi378decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var viewRecordsTests = []struct {
	fmd            []byte
	expectedResult []ViewRecord
	expectedErr    error
}{
	{invalidShortFMD, nil, ErrInvalidFMD},
	{valid2ByteFMD, valid2ByteViewRecords, nil},
}

func TestViewRecords(t *testing.T) {
	assert := assert.New(t)
	for _, test := range viewRecordsTests {
		recordHeaders, err := RecordHeaders(test.fmd)
		assert.EqualValues(test.expectedErr, err)
		// if we're doing a test of shorter-than-valid fmds
		var res []ViewRecord
		if len(test.fmd) < 27 {
			// don't cut the slice
			res, err = ViewRecords(test.fmd, recordHeaders.Views)
		} else { // chop the header off
			res, err = ViewRecords(test.fmd[26:], recordHeaders.Views)
		}
		assert.EqualValues(test.expectedResult, res)
		assert.EqualValues(test.expectedErr, err)
	}
}
