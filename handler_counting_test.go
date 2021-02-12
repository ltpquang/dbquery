package dbquery

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestCountingHandler_Target(t *testing.T) {
	handler := countingHandler{}
	assert.Equal(t, typeCount, handler.Target())
}

func TestCountingHandler_Apply(t *testing.T) {
	var integer int
	var integer8 int8
	var integer16 int16
	var integer32 int32
	var integer64 int64
	var uinteger uint
	var uinteger8 uint8
	var uinteger16 uint16
	var uinteger32 uint32
	var uinteger64 uint64
	var aList = make([]interface{}, 0)
	var aMap = make(map[string]string)

	tests := []struct {
		input         interface{}
		expectedError error
	}{
		{
			input:         nil,
			expectedError: nil,
		},
		{
			input:         &integer,
			expectedError: ErrCastingError,
		},
		{
			input:         &integer8,
			expectedError: ErrCastingError,
		},
		{
			input:         &integer16,
			expectedError: ErrCastingError,
		},
		{
			input:         &integer32,
			expectedError: ErrCastingError,
		},
		{
			input:         &integer64,
			expectedError: nil,
		},
		{
			input:         &uinteger,
			expectedError: ErrCastingError,
		},
		{
			input:         &uinteger8,
			expectedError: ErrCastingError,
		},
		{
			input:         &uinteger16,
			expectedError: ErrCastingError,
		},
		{
			input:         &uinteger32,
			expectedError: ErrCastingError,
		},
		{
			input:         &uinteger64,
			expectedError: ErrCastingError,
		},
		{
			input:         integer,
			expectedError: ErrCastingError,
		},
		{
			input:         integer8,
			expectedError: ErrCastingError,
		},
		{
			input:         integer16,
			expectedError: ErrCastingError,
		},
		{
			input:         integer32,
			expectedError: ErrCastingError,
		},
		{
			input:         integer64,
			expectedError: ErrCastingError,
		},
		{
			input:         uinteger,
			expectedError: ErrCastingError,
		},
		{
			input:         uinteger8,
			expectedError: ErrCastingError,
		},
		{
			input:         uinteger16,
			expectedError: ErrCastingError,
		},
		{
			input:         uinteger32,
			expectedError: ErrCastingError,
		},
		{
			input:         uinteger64,
			expectedError: ErrCastingError,
		},
		{
			input:         aList,
			expectedError: ErrCastingError,
		},
		{
			input:         aMap,
			expectedError: ErrCastingError,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			handler := &countingHandler{}
			db, _ := gorm.Open(sqlite.Open(path.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
			db, err := handler.Apply(db, test.input)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, test.expectedError))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
