package main

import (
	"fmt"
	"time"

	"github.com/cryptrunner49/nanotrace/internal/tracer"
	"github.com/cryptrunner49/nanotrace/pkg/nanotrace"
)

func main() {
	fmt.Println("Starting NanoTrace...")

	// Start profiling
	profiler := nanotrace.NewProfiler()
	profiler.Start()
	defer profiler.Stop()

	// Start execution tracing
	tracer := tracer.NewTracer()
	tracer.Start()
	defer tracer.Stop()

	// Simulate workload
	simulateWorkload()

	fmt.Println("Profiling and tracing completed.")
}

func simulateWorkload() {
	fmt.Println("Simulating workload...")
	time.Sleep(200 * time.Millisecond) // Simulated function execution time
}
