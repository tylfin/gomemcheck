package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gomemcheck/internal/test"
)

func TestCheckMemory(t *testing.T) {
	test.MemLeak()
	leak := Check(t)
	assert.Equal(t, leak, true)
}
