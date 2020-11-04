package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gomemcheck/internal/test"
)

// MockTB records if Verify successfully calls the Fail() function when a memleak is present
type MockTB struct {
	*testing.T
	calledFail bool
}

func (m *MockTB) Fail() {
	m.calledFail = true
	return
}

func (m *MockTB) Log(args ...interface{}) {
	fmt.Println(args...)
}

func TestVerifyNoLeak(t *testing.T) {
	t.Skip("¯\\_(ツ)_/¯")
	// TestVerifyNoLeak should always pass given that MemSafe(...) properly cleans itself up
	m := &MockTB{}
	test.MemSafe()
	Verify(m)

	assert.Equal(t, m.calledFail, false)
}

func TestVerifyLeak(t *testing.T) {
	m := &MockTB{}
	test.MemLeak()
	Verify(m)

	assert.Equal(t, m.calledFail, true)
}
