package test

import (
	"time"
)

const tickTime = time.Second

// MemLeak creates a memory leak which uses the time.Tick function which is documented as: cannot be recovered by the
// garbage collector; it "leaks".
func MemLeak() {
	select {
	case <-time.Tick(tickTime):
		return
	}
}

// MemSafe creates a NewTicker and stops it after the function returns. This should be cleaned up by the garbage
// collector.
func MemSafe() {
	t := time.NewTicker(tickTime)
	// defer t.Stop()

	select {
	case <-t.C:
		// Setting the reference count to 0 should accelerate the time to garbage collection.
		// t = nil

		return
	}
}
