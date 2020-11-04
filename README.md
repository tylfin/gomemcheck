# Go Memory Check

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
![Go](https://github.com/tylfin/gomemcheck/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tylfin/gomemcheck)](https://goreportcard.com/report/github.com/tylfin/gomemcheck)
![CodeQL](https://github.com/tylfin/gomemcheck/workflows/CodeQL/badge.svg)

Go memory leak detector to help avoid memory leaks.

## Example

To detect a memory leak defer the `Verify` function call in your test:

```go
import "github.com/tylfin/gomemcheck"

func TestLeak(t *testing.T) {
    defer gomemcheck.Verify(t)
    MemLeak()
}
```

When the test suite runs a "Memory leak detected:" failure will appear with additional information:

```go
$ go test
--- FAIL: TestLeak (1.00s)
    code.go:LINE: Memory Leak Detected:
    code.go:LINE: 1: 32 [1: 32] @ 0x1058871 0x10591dd 0x1023ccc 0x1023c89 0x1023c79 0x1023f2d 0x106f261
        #       0x1023c88+time.resetTimer+0x48  :
        #       0x1023c78+runtime.scavengeSleep+0x38   /go/.../src/runtime/mgcscavenge.go:237
        #       0x1023f2c+runtime.bgscavenge+0x1ec     /go/.../src/runtime/mgcscavenge.go:366

--- FAIL: TestLeak (1.01s)
...
```

## Stability

This is currently an experimental package under development based-off [goleak](https://github.com/uber-go/goleak) by
uber. The API is very likely to change.

Note that this is currently **not** working, see `main_test.TestVerifyNoLeak` for more information on the issue.
