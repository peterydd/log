package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHello - verify reference when logging from a regular function
func TestHello(t *testing.T) {
	fields := Hello("test")

	assert.Equal(t, "Hi, test. Welcome!", fields)
}
