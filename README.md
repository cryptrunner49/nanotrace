# NanoTrace 🚀

A lightweight performance profiler and tracer for Golang to analyze function execution times efficiently. ⚡

## Overview 📊

NanoTrace provides an easy-to-use toolkit to profile and trace your Golang applications. It helps you:

- **Measure Function Execution Times** ⏱️
- **Identify Performance Bottlenecks** 🔍
- **Optimize Application Speed** 🚀

## Features ✨

- **Minimal Overhead**: Lightweight and efficient with near-zero impact on performance.
- **Simple API**: Easy-to-use functions to start and stop profiling.
- **Real-time Insights**: View execution times directly in the console.
- **Customizable**: Flexible settings for different tracing levels.

## Installation 📥

Ensure you have Go installed, then install NanoTrace using:

```bash
go get github.com/cryptrunner49/nanotrace
```

## Usage 💻

Here’s a quick example to get started:

```go
package main

import (
    "fmt"
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

func someFunction() int {
    start := time.Now()
    defer nanotrace.Track("someFunction", start)
    time.Sleep(100 * time.Millisecond) // Simulating work
    return 42
}
```

## Contributing 🤝

We welcome contributions! Check out our [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines.

## License 📄

This project is licensed under the MIT License.
