package ansi378decoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var recordHeaderTests = []struct {
	fmd            []byte
	expectedResult *RecordHeader
	expectedErr    error
}{
	{invalidShortFMD, nil, ErrInvalidFMD},
	{valid2ByteFMD, &validHeader, nil},
	{valid6ByteFMD, &validHeader, nil},
}

func TestRecordHeaders(t *testing.T) {

	assert := assert.New(t)
	for _, test := range recordHeaderTests {
		res, err := RecordHeaders(test.fmd)
		assert.EqualValues(test.expectedResult, res)
		assert.EqualValues(test.expectedErr, err)
	}

}
