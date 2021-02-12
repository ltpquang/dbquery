package dbquery

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPreloadHandler_Target(t *testing.T) {
	handler := preloadHandler{}
	assert.Equal(t, typePreload, handler.Target())
}

func TestPreloadHandler_Apply(t *testing.T) {
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
			expectedError: nil,
		},
		{
			input: &QueryKeyValue{
				key:   "preload",
				value: "a string",
			},
			expectedError: ErrCastingError,
		},
		{
			input: &QueryKeyValue{
				key:   "preload",
				value: make(map[string]string),
			},
			expectedError: ErrCastingError,
		},
		{
			input: &QueryKeyValue{
				key:   "preload",
				value: make([]interface{}, 0),
			},
			expectedError: nil,
		},
		{
			input: &QueryKeyValue{
				key:   "preload",
				value: make([]string, 0),
			},
			expectedError: ErrCastingError,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			handler := &preloadHandler{}
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
