package nanotrace

import (
	"fmt"
	"time"
)

// Profiler struct holds the start time
type Profiler struct {
	startTime time.Time
}

// NewProfiler initializes a new Profiler
func NewProfiler() *Profiler {
	return &Profiler{}
}

// Start begins CPU profiling
func (p *Profiler) Start() {
	p.startTime = time.Now()
	fmt.Println("CPU profiling started...")
}

// Stop stops profiling and prints elapsed time
func (p *Profiler) Stop() {
	elapsed := time.Since(p.startTime)
	fmt.Printf("CPU profiling completed. Total execution time: %v\n", elapsed)
}
