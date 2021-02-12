package dbquery

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInListHandler_Target(t *testing.T) {
	handler := inListHandler{}
	assert.Equal(t, typeInList, handler.Target())
}

func TestInListHandler_Apply(t *testing.T) {
	tests := []struct {
		input         interface{}
		expectedError error
	}{
		{
			input:         QueryKeyValue{},
			expectedError: ErrCastingError,
		},
		{
			input:         &QueryKeyValue{},
			expectedError: ErrCastingError,
		},
		{
			input: &QueryKeyValue{
				key:   "inlist",
				value: "a string",
			},
			expectedError: ErrCastingError,
		},
		{
			input: &QueryKeyValue{
				key:   "inlist",
				value: make(map[string]string),
			},
			expectedError: ErrCastingError,
		},
		{
			input: &QueryKeyValue{
				key:   "inlist",
				value: make([]interface{}, 0),
			},
			expectedError: nil,
		},
		{
			input: &QueryKeyValue{
				key:   "inlist",
				value: make([]string, 0),
			},
			expectedError: ErrCastingError,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			handler := &inListHandler{}
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
