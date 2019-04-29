package ansi378decoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var viewRecordsTests = []struct {
	fmd            []byte
	expectedResult *[]ViewRecord
	expectedErr    error
}{
	{invalidShortFMD, nil, ErrInvalidFMD},
	{valid2ByteFMD, valid2ByteViewRecords, nil},
}

func TestViewRecords(t *testing.T) {
	assert := assert.New(t)
	for _, test := range viewRecordsTests {
		recordHeaders, err := RecordHeaders(test.fmd)
		if err != nil {
			assert.EqualValues(test.expectedErr, err)
		}
		if recordHeaders != nil {
			res, err := ViewRecords(test.fmd[26:], recordHeaders.Views)
			if err == nil {
				assert.EqualValues(test.expectedResult, res)
			}
		}
		assert.EqualValues(test.expectedErr, err)
	}
}

