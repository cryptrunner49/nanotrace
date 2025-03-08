package nanotrace

import (
    "os"
    "testing"
)

func TestProfiler(t *testing.T) {
    profiler := NewProfiler()
    err := profiler.Start()
    if err != nil {
        t.Fatalf("Failed to start profiler: %v", err)
    }
    profiler.Stop()

    if _, err := os.Stat("cpu_profile.prof"); os.IsNotExist(err) {
        t.Errorf("CPU profile file was not created")
    } else {
        os.Remove("cpu_profile.prof")
    }
}
