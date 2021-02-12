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

func TestMatchingHandler_Target(t *testing.T) {
	handler := matchingHandler{}
	assert.Equal(t, typeMatch, handler.Target())
}

func TestMatchingHandler_Apply(t *testing.T) {
	tests := []struct {
		input         interface{}
		expectedError error
	}{
		{
			input:         ComparingObject{},
			expectedError: ErrCastingError,
		},
		{
			input:         &ComparingObject{},
			expectedError: nil,
		},
		{
			input:         &QueryKeyValue{},
			expectedError: ErrCastingError,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			handler := &matchingHandler{}
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
