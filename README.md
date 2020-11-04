# Go Memory Check

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
![Go](https://github.com/tylfin/gomemcheck/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tylfin/gomemcheck)](https://goreportcard.com/report/github.com/tylfin/gomemcheck)
![CodeQL](https://github.com/tylfin/gomemcheck/workflows/CodeQL/badge.svg)

Go memory leak detector to help avoid memory leaks. This is currently an experimental package under development based
off similar goroutine leak packages.

## Example

To detect a memory leak simply defer the `Verify` call in your test like so:

```go
import "github.com/tylfin/gomemcheck"

func TestLeak(t *testing.T) {
    defer gomemcheck.Verify(m)
    MemLeak()
}
```

When you run your test you'll now see the "Memory leak detected:" warning with additional information:

```go
$ go test
--- FAIL: TestLeak (1.00s)
...
Memory leak detected: [Additional information here]
FAIL
...
```
