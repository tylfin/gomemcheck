package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var IgnoreEntryRecords = []string{
	// By default, Check should ignore anything on the heap that it's allocated
	"gomemcheck/internal",
	"gomemcheck.Verify",
	"runtime/pprof",
	"runtime.systemstack",
	"testing.tRunner",
}

// StackRecord stores the information for each goroutine's heap profile
type StackRecord struct {
	Info         string
	FrameRecords []FrameRecord
}

type FrameRecord struct {
	// PC is the program counter for the location in this frame.
	PC   string
	Name string
	// Entry point program counter for the function
	Entry string
	// File and Line are the file name and line number of the location in this frame.F
	File string
	Line string
}

func Check(t testing.TB) bool {
	stackRecords, err := getStackRecords()
	if err != nil {
		t.Logf("Gomemcheck failure: %s", err)
		return false
	}

	logInfo := ""

	if len(stackRecords) > 0 {
		for _, stackRecord := range stackRecords {
			if len(stackRecord.FrameRecords) < 1 {
				continue
			}

			logInfo += stackRecord.Info + "\n"
			for _, frame := range stackRecord.FrameRecords[1:] {
				logInfo += fmt.Sprintf("%s\t%s+%s\t%s:%s\n", frame.PC, frame.Name, frame.Entry, frame.File, frame.Line)
			}

			logInfo += "\n\n"
		}

		t.Log("Memory Leak Detected:")
		t.Log(logInfo)
		return true
	}

	return false
}

func getStackRecords() ([]StackRecord, error) {
	// Run the garbage collector twice and set MemProfileRate to get latest, and maximum
	// amount of information about the current heap allocations
	debug.SetGCPercent(100)
	runtime.GC()
	runtime.GC()
	debug.FreeOSMemory()
	runtime.MemProfileRate = 1

	stackRecords := []StackRecord{}
	for _, profile := range pprof.Profiles() {
		if profile.Name() != "heap" {
			continue
		}

		// Re-direct everything from profile.WriteTo to our reader
		reader, writer := io.Pipe()
		go func() {
			// Use WriteTo here because the profiles have some important internal locking mechanisms
			// we want to take advantage of without reimplementing
			profile.WriteTo(writer, 1)
			writer.Close()
		}()

		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, errors.Wrap(err, "could not properly read the heap profile")
		}
		reader.Close()

		stacks := strings.Split(string(b)[1:], "# runtime.MemStats")[0]
		for _, stack := range strings.Split(stacks, "\n\n") {
			stackRecord := StackRecord{Info: strings.Split(stack, "\n")[0], FrameRecords: []FrameRecord{}}

			if !(len(strings.Split(stack, "\n")) > 1) {
				continue
			}

			for _, frame := range strings.Split(stack[1:], "\n") {
				frameRecord := FrameRecord{}

				for c, v := range strings.Split(frame, "\t") {
					switch c {
					case 0:
						frameRecord.PC = v
					case 1:
						frameRecord.Name = v
					case 2:
						frameRecord.Entry = v

						// Skip over any FrameRecords that are contained in the IgnoreFileRecords array
						for _, ignoreEntryRecord := range IgnoreEntryRecords {
							if strings.Contains(frameRecord.Entry, ignoreEntryRecord) {
								goto skip
							}
						}
					case 3:
						if len(strings.Split(v, ":")) < 2 {
							continue
						}
						frameRecord.File = strings.Split(v, ":")[0]
						frameRecord.Line = strings.Split(v, ":")[1]
					}
				}

				stackRecord.FrameRecords = append(stackRecord.FrameRecords, frameRecord)
			skip:
			}

			if stackRecord.Info == "" || len(stackRecord.FrameRecords) <= 1 {
				continue
			}

			stackRecords = append(stackRecords, stackRecord)
		}
	}

	return stackRecords, nil
}
