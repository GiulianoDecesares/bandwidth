# Bandwidth

Bandwidth is a lightweight Go package that provides a C# TimeSpan-like abstraction for representing, converting, and formatting network bandwidth (data transfer rate). With support for multiple units (bits, bytes, kilobits, megabits, etc.), this package makes it easy to work with bandwidth values in a human-readable way.

## Features

- **Create Bandwidth values:** Easily instantiate a Bandwidth value using your desired unit.
- **Unit Conversion:** Convert bandwidth values between various units.
- **Human-Readable Output:** Format bandwidth values as strings with a user-specified unit.
- **Clean API:** Simple functions and a clean API make integration straightforward.

## Installation

Use `go get` to install the module:

```bash
go get github.com/yourusername/bandwidth
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/bandwidth"
)

func main() {
	bw := bandwidth.New(5, bandwidth.MegabytePerSecond)
	
	fmt.Println("Bandwidth:", bw.String(bandwidth.MegabytePerSecond))
	
	// Convert the bandwidth to KB/s.
	valueInKB := bw.To(bandwidth.KilobytePerSecond)
	fmt.Printf("Equivalent to: %.2f KB/s\n", valueInKB)
}
```
