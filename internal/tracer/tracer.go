package tracer

import (
    "os"
    "runtime/trace"
)

var traceFile *os.File

// Start initializes tracing and writes to the specified file.
func Start(filename string) error {
    var err error
    traceFile, err = os.Create(filename)
    if err != nil {
        return err
    }
    return trace.Start(traceFile)
}

// Stop ends the tracing session.
func Stop() {
    trace.Stop()
    if traceFile != nil {
        traceFile.Close()
    }
}
