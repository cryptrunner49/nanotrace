
# NanoTrace Documentation

## NanoTrace - A Simple Execution Time Tracing Library

`nanotrace` is a lightweight Go library designed to measure and report the execution time of code blocks. It provides a simple API to start and stop tracing, outputting the results to a specified `io.Writer`. This library is ideal for debugging, performance monitoring, or logging execution durations in Go applications.

### Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Custom Output](#custom-output)
  - [Error Handling](#error-handling)
- [API Documentation](#api-documentation)
  - [Tracer Struct](#tracer-struct)
  - [NewTracer](#newtracer)
  - [Start](#start)
  - [Stop](#stop)
- [Testing](#testing)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

To use the `nanotrace` library in your Go project, install it via `go get`:

```bash
go get github.com/cryptrunner49/nanotrace
```

Ensure your Go module is initialized (`go mod init`) if it isn’t already.

---

## Quick Start

Here’s a minimal example to get started with `nanotrace`:

```go
package main

import (
 "fmt"
 "os"
 "time"

 "github.com/cryptrunner49/nanotrace/tracer"
)

func main() {
 // Create a new tracer with stdout as the output
 tracer := tracer.NewTracer(os.Stdout)
 
 // Start tracing
 tracer.Start()
 
 // Simulate some work
 time.Sleep(200 * time.Millisecond)
 
 // Stop tracing and report the time
 defer tracer.Stop()
}
```

**Output:**

```text
Tracing started...
Tracing completed. Execution time: 200.123ms
```

---

## Usage

### Basic Usage

The `Tracer` type allows you to measure the time between `Start()` and `Stop()` calls. It writes messages to the provided `io.Writer` (e.g., `os.Stdout`).

```go
tr := tracer.NewTracer(os.Stdout)
tr.Start()
// Your code here
tr.Stop()
```

### Custom Output

You can direct tracing output to any `io.Writer`, such as a file, buffer, or logger:

```go
package main

import (
 "os"
 "github.com/cryptrunner49/nanotrace/tracer"
)

func main() {
 // Write to a file instead of stdout
 file, _ := os.Create("trace.log")
 defer file.Close()
 
 tr := tracer.NewTracer(file)
 tr.Start()
 // Do work
 tr.Stop()
}
```

### Error Handling

The library is designed to be safe and idempotent:

- Calling `Start()` multiple times only starts tracing once.
- Calling `Stop()` before `Start()` does nothing.
- If the `io.Writer` is `nil`, no output is written.

No explicit error handling is required, but you should ensure the `io.Writer` is valid if output is critical.

---

## API Documentation

### Tracer Struct

```go
type Tracer struct {
 startTime time.Time // Time when tracing started
 started   bool      // Indicates if tracing has started
 out       io.Writer // Output destination for tracing messages
}
```

The `Tracer` struct tracks the tracing state and output destination. It should be instantiated using `NewTracer`.

### NewTracer

```go
func NewTracer(out io.Writer) *Tracer
```

**Description:** Creates a new `Tracer` instance with the specified output writer.

- **Parameters:**
  - `out io.Writer`: The destination for tracing messages (e.g., `os.Stdout`, a file, or `nil` for no output).
- **Returns:** A pointer to a new `Tracer` instance.

**Example:**

```go
tr := tracer.NewTracer(os.Stdout)
```

### Start

```go
func (t *Tracer) Start()
```

**Description:** Begins tracing by recording the current time. If tracing has already started, this does nothing.

- **Receiver:** `*Tracer`
- **Side Effects:** Writes "Tracing started..." to the output writer if not `nil`.

**Example:**

```go
tr.Start()
```

### Stop

```go
func (t *Tracer) Stop()
```

**Description:** Ends tracing and reports the elapsed time since `Start()` was called. If tracing hasn’t started, this does nothing.

- **Receiver:** `*Tracer`
- **Side Effects:** Writes "Tracing completed. Execution time: `<duration>`" to the output writer if not `nil`.

**Example:**

```go
tr.Stop()
```

---

## Testing

The library includes a comprehensive test suite in `tracer/tracer_test.go`. To run the tests:

```bash
cd tracer
go test -v
```

The tests cover:

- Normal tracing flow with expected output and timing.
- Calling `Stop()` before `Start()` (no-op).
- Multiple `Start()` calls (idempotent behavior).

---

## Examples

### Example 1: Basic Tracing

```go
package main

import (
 "fmt"
 "os"
 "time"
 "github.com/cryptrunner49/nanotrace/tracer"
)

func main() {
 tr := tracer.NewTracer(os.Stdout)
 tr.Start()
 time.Sleep(100 * time.Millisecond)
 tr.Stop()
}
```

### Example 2: Tracing to a Buffer

```go
package main

import (
 "strings"
 "github.com/cryptrunner49/nanotrace/tracer"
)

func main() {
 var buf strings.Builder
 tr := tracer.NewTracer(&buf)
 tr.Start()
 // Do work
 tr.Stop()
 println(buf.String())
}
```

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please include tests for new features or bug fixes.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
