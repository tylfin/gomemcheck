package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMemory(t *testing.T) {
	assert.NotPanics(t, func() { CheckMemory() })
}
