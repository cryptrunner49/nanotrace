# ⚡ NanoTrace 🚀

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
- **Fully Customizable**: Tailor the output writer to your needs. If out is nil, no output is generated.

## Quick Start 🚀

1. **Install via Go Modules:**

   ```bash
   go get github.com/cryptrunner49/nanotrace
   ```  

2. **Basic Usage Example:**

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

func simulateWorkload() {
 fmt.Println("Simulating workload...")
 time.Sleep(200 * time.Millisecond) // Simulated function execution time
}
```

## Contributing 🤝

We welcome contributions! Check out our [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines.

## License 📜

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.
