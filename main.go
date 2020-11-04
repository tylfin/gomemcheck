package main

import (
	"runtime"
	"testing"

	"code.cloudfoundry.org/bytefmt"
)

// CheckMemory prints out the size of the allocated heap objects, the cumulative size of allocated for heap objects,
// the total bytes of memory obtained from the OS, and the number of completed GC cycles. This is just to start
// experimenting with the runtime library.
func CheckMemory() (alloc, totalAlloc, sys string, gc uint32) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	alloc = bytefmt.ByteSize(m.Alloc)
	totalAlloc = bytefmt.ByteSize(m.TotalAlloc)
	sys = bytefmt.ByteSize(m.Sys)
	gc = m.NumGC
	return
}

// Verify ensures that all objects that garbage collection successfully cleans up any left-over heap objects.
func Verify(t testing.TB) {
	return
}

func main() {
	CheckMemory()
}
