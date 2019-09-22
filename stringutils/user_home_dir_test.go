package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHomeDir(t *testing.T) {

	// Arrange && Act.

	result := UserHomeDir()

	// Assert.

	assert.Equal(t, "/Users/l0cke", result)
}
