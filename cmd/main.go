package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cryptrunner49/nanotrace/tracer"
)

func main() {
	fmt.Println("Starting NanoTrace...")

	// Start execution tracing with stdout as the output destination
	tracer := tracer.NewTracer(os.Stdout)
	tracer.Start()
	defer tracer.Stop()

	// Simulate workload
	simulateWorkload()
}

func simulateWorkload() {
	fmt.Println("Simulating workload...")
	time.Sleep(200 * time.Millisecond) // Simulated function execution time
}
