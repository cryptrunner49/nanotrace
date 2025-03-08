package tracer

import (
	"strings"
	"testing"
	"time"
)

func TestTracer(t *testing.T) {
	t.Run("NormalFlow", func(t *testing.T) {
		var buf strings.Builder // Use strings.Builder for efficiency
		tracer := NewTracer(&buf)

		// Run the tracer
		tracer.Start()
		time.Sleep(100 * time.Millisecond) // Simulate work
		tracer.Stop()

		// Verify output
		lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("Expected 2 lines of output, got %d:\n%s", len(lines), buf.String())
		}
		if lines[0] != "Tracing started..." {
			t.Errorf("Expected first line 'Tracing started...', got: %s", lines[0])
		}
		if !strings.HasPrefix(lines[1], "Tracing completed. Execution time: ") {
			t.Errorf("Expected second line to start with 'Tracing completed. Execution time: ', got: %s", lines[1])
		}

		// Check execution time (should be ~100ms)
		timeStr := strings.TrimPrefix(lines[1], "Tracing completed. Execution time: ")
		elapsed, err := time.ParseDuration(timeStr)
		if err != nil {
			t.Errorf("Failed to parse execution time: %v", err)
		}
		if elapsed < 90*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("Expected execution time ~100ms, got: %v", elapsed)
		}
	})

	t.Run("StopBeforeStart", func(t *testing.T) {
		var buf strings.Builder
		tracer := NewTracer(&buf)
		tracer.Stop() // Should do nothing
		if buf.Len() != 0 {
			t.Errorf("Expected no output when Stop called before Start, got: %s", buf.String())
		}
	})

	t.Run("MultipleStartCalls", func(t *testing.T) {
		var buf strings.Builder
		tracer := NewTracer(&buf)

		tracer.Start()
		time.Sleep(50 * time.Millisecond)
		tracer.Start() // Should be ignored
		time.Sleep(50 * time.Millisecond)
		tracer.Stop()

		lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("Expected 2 lines of output, got %d:\n%s", len(lines), buf.String())
		}

		// Check execution time (should be ~100ms, not 50ms)
		timeStr := strings.TrimPrefix(lines[1], "Tracing completed. Execution time: ")
		elapsed, err := time.ParseDuration(timeStr)
		if err != nil {
			t.Errorf("Failed to parse execution time: %v", err)
		}
		if elapsed < 90*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("Expected execution time ~100ms (ignoring second Start), got: %v", elapsed)
		}
	})
}
