package main

import (
	"fmt"
	"gomemcheck/test"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestCheckMemory(t *testing.T) {
	alloc, totalAlloc, sys, gc := CheckMemory()
	assert.Equal(t, gc >= 0, true)
	fmt.Println(alloc, totalAlloc, sys)
}

func TestVerifyNoLeak(t *testing.T) {
	// TestVerifyNoLeak should always pass given that MemSafe(...) properly cleans itself up
	m := &MockTB{}
	test.MemSafe()
	Verify(m)

	assert.Equal(t, m.calledFail, false)
}

func TestVerifyLeak(t *testing.T) {
	t.Skip("Not implemented yet")
	m := &MockTB{}
	test.MemLeak()
	Verify(m)

	assert.Equal(t, m.calledFail, true)
}
