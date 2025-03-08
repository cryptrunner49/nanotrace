package nanotrace

import (
    "os"
    "runtime/pprof"
)

type Profiler struct {
    cpuProfileFile *os.File
}

// NewProfiler creates a new Profiler instance.
func NewProfiler() *Profiler {
    return &Profiler{}
}

// Start begins CPU profiling and writes to 'cpu_profile.prof'.
func (p *Profiler) Start() error {
    var err error
    p.cpuProfileFile, err = os.Create("cpu_profile.prof")
    if err != nil {
        return err
    }
    return pprof.StartCPUProfile(p.cpuProfileFile)
}

// Stop ends the CPU profiling session.
func (p *Profiler) Stop() {
    pprof.StopCPUProfile()
    if p.cpuProfileFile != nil {
        p.cpuProfileFile.Close()
    }
}
