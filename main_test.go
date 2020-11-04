package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMemory(t *testing.T) {
	alloc, totalAlloc, sys, gc := CheckMemory()
	assert.Equal(t, gc >= 0, true)
	fmt.Println(alloc, totalAlloc, sys)
	t.Fail()
}
