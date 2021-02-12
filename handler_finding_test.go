package dbquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindingHandler_Target(t *testing.T) {
	handler := findingHandler{}
	assert.Equal(t, typeFind, handler.Target())
}
