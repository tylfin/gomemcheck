package internal

import (
	"gomemcheck/internal/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMemory(t *testing.T) {
	test.MemLeak()
	leak := Check(t)
	assert.Equal(t, leak, true)
}
