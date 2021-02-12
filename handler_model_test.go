package dbquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelHandler_Target(t *testing.T) {
	handler := modelHandler{}
	assert.Equal(t, typeModel, handler.Target())
}
