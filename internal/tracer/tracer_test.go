package tracer

import (
    "os"
    "testing"
)

func TestTracer(t *testing.T) {
    filename := "test_trace.out"
    err := Start(filename)
    if err != nil {
        t.Fatalf("Failed to start tracer: %v", err)
    }
    Stop()

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        t.Errorf("Trace file %s was not created", filename)
    } else {
        os.Remove(filename)
    }
}
