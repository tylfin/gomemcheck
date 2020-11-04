package main

import (
	"gomemcheck/internal"
	"testing"
)

// Verify ensures that all objects that garbage collection successfully cleans up any left-over heap objects.
func Verify(t testing.TB) {
	if leak := internal.Check(t); leak {
		t.Fail()
	}
}
