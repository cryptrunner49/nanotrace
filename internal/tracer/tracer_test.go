package tracer

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestTracer(t *testing.T) {
	var buf bytes.Buffer

	// Redirect stdout to the buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
		w.Close()
	}()

	// Capture the output in a separate goroutine
	go func() {
		io.Copy(&buf, r)
		r.Close()
	}()

	// Create and use the tracer
	tracer := NewTracer()
	tracer.Start()
	time.Sleep(100 * time.Millisecond)
	tracer.Stop()

	// Restore stdout and check the output
	w.Close()
	output := buf.String()
	if !contains(output, "Tracing started...") || !contains(output, "Tracing completed.") {
		t.Errorf("Expected tracing output, got:\n%s", output)
	}
}

// Helper function to check if the output contains the given substring
func contains(output, substr string) bool {
	return bytes.Contains([]byte(output), []byte(substr))
}
