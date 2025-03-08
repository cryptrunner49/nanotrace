package tracer

import (
	"fmt"
	"io"
	"time"
)

// Tracer measures and reports the execution time of a process.
type Tracer struct {
	startTime time.Time // Time when tracing started
	started   bool      // Indicates if tracing has started
	out       io.Writer // Output destination for tracing messages
}

// NewTracer initializes a new Tracer with a custom output writer.
// If out is nil, no output will be written.
func NewTracer(out io.Writer) *Tracer {
	return &Tracer{
		out: out,
	}
}

// Start begins execution tracing by recording the current time.
// It does nothing if tracing has already started.
func (t *Tracer) Start() {
	if t.started {
		return // Ignore if already started
	}
	t.startTime = time.Now()
	t.started = true
	if t.out != nil {
		fmt.Fprintln(t.out, "Tracing started...")
	}
}

// Stop ends tracing and reports the elapsed time since Start was called.
// It does nothing if tracing hasn’t started.
func (t *Tracer) Stop() {
	if !t.started {
		return // Ignore if not started
	}
	elapsed := time.Since(t.startTime)
	t.started = false
	if t.out != nil {
		fmt.Fprintf(t.out, "Tracing completed. Execution time: %v\n", elapsed)
	}
}
