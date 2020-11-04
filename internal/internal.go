package internal

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"testing"

	"code.cloudfoundry.org/bytefmt"
)

const goheap17 = "go1.7 heap dump"

func Check(t testing.TB) {
	dumpPath, err := writeDump(t)
	if err != nil {
		t.Logf("Gomemcheck: Could not write heap dump file: %s", err)
		return
	}

	b, err := readDump(dumpPath)
	if err != nil {
		t.Logf("Gomemcheck: Could not read heap dump: %s", err)
		return
	}

	switch string(b[:len(goheap17)]) {
	case goheap17:
	default:
		t.Logf("Gomemcheck: heap dump format not implemented: %s", b[:len(goheap17)])
		return
	}
}

func readDump(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func writeDump(t testing.TB) (string, error) {
	tmpDir := t.TempDir()
	dump := path.Join(tmpDir, "dump")

	f, err := os.Create(dump)
	if err != nil {
		return "", err
	}
	defer f.Close()

	debug.WriteHeapDump(f.Fd())
	return dump, nil
}

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
