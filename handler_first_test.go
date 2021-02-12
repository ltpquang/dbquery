package dbquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstHandler_Target(t *testing.T) {
	handler := findFirstHandler{}
	assert.Equal(t, typeFirst, handler.Target())
}
