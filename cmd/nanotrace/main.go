package main

import (
    "fmt"
    "time"

    "github.com/cryptrunner49/nanotrace/internal/tracer"
    "github.com/cryptrunner49/nanotrace/pkg/nanotrace"
)

func main() {
    // Start tracing
    traceFile := "trace.out"
    if err := tracer.Start(traceFile); err != nil {
        fmt.Printf("Failed to start tracer: %v\n", err)
        return
    }
    defer tracer.Stop()

    // Start profiling
    profiler := nanotrace.NewProfiler()
    profiler.Start()
    defer profiler.Stop()

    // Simulate application workload
    simulateWorkload()

    fmt.Println("Profiling and tracing completed. Analyze 'cpu_profile.prof' and 'trace.out' for results.")
}

func simulateWorkload() {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
    }
}
