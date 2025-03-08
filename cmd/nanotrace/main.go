package main

import (
	"fmt"
	"time"

	"github.com/cryptrunner49/nanotrace/internal/tracer"
)

func main() {
	fmt.Println("Starting NanoTrace...")

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
