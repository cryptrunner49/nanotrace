package tracer

import (
	"fmt"
	"time"
)

// Tracer struct holds the start time
type Tracer struct {
	startTime time.Time
}

// NewTracer initializes a new Tracer
func NewTracer() *Tracer {
	return &Tracer{}
}

// Start begins execution tracing
func (t *Tracer) Start() {
	t.startTime = time.Now()
	fmt.Println("Tracing started...")
}

// Stop stops tracing and prints execution time
func (t *Tracer) Stop() {
	elapsed := time.Since(t.startTime)
	fmt.Printf("Tracing completed. Execution time: %v\n", elapsed)
}
