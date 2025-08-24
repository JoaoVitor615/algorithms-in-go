# 🛠️ Package `pkg` - Utility Functions

This package provides a collection of generic utility functions designed to reduce code duplication and improve maintainability across the algorithms-in-go project.

## 📁 Structure

```
pkg/
├── README.md           # This documentation
├── input.go           # User input utilities
├── format.go          # Formatting and display utilities
├── generator.go       # Random data generation utilities
├── validator.go       # Data validation utilities
└── performance.go     # Performance analysis utilities
```

## 🔧 Modules

### 📥 **Input Module** (`input.go`)

Provides utilities for reading and validating user input.

**Key Types:**
- `InputReader` - Main input handling struct

**Key Functions:**
```go
// Create new input reader
reader := pkg.NewInputReader()

// Read string input
text := reader.ReadString("Enter text: ")

// Read integer with validation
num, err := reader.ReadInt("Enter number: ", 1, 100)

// Read integer or return -1 for invalid
num := reader.ReadIntOrDefault("Enter number: ", 1, 100)

// Read sequence of numbers
numbers := reader.ReadNumbers()

// Read yes/no response
confirmed := reader.ReadYesNo("Continue? (y/n): ")
```

### 🎨 **Format Module** (`format.go`)

Provides utilities for formatting numbers and displaying data.

**Key Functions:**
```go
// Format number with thousands separators
formatted := pkg.FormatNumber(1234567) // "1,234,567"

// Print array in bracket format
pkg.PrintSlice([]int{1, 2, 3}) // [1, 2, 3]

// Print partial array
pkg.PrintSlicePartial([]int{1, 2, 3, 4, 5}, 3) // [1, 2, 3, ... (and more)]

// Print formatted headers
pkg.PrintHeader("BENCHMARK RESULTS")
pkg.PrintSubHeader("Quick Sort Test")
pkg.PrintSeparator(50)
```

### 🎲 **Generator Module** (`generator.go`)

Provides utilities for generating random test data.

**Key Types:**
- `RandomGenerator` - Main random data generator

**Key Functions:**
```go
// Create generator with current time seed
gen := pkg.NewRandomGenerator()

// Create generator with specific seed
gen := pkg.NewRandomGeneratorWithSeed(12345)

// Generate random integer slice
numbers := gen.GenerateIntSlice(100, 1, 1000) // 100 numbers between 1-1000

// Generate with default range (1-1000)
numbers := gen.GenerateIntSliceDefault(100)

// Generate sorted slice
sorted := gen.GenerateSortedSlice(50, 1, 100)

// Generate reverse sorted slice
reverse := gen.GenerateReverseSortedSlice(50, 1, 100)

// Shuffle existing slice
gen.ShuffleSlice(numbers)
```

### ✅ **Validator Module** (`validator.go`)

Provides utilities for data validation and analysis.

**Key Functions:**
```go
// Check if slice is sorted
isSorted := pkg.IsSortedSlice([]int{1, 2, 3, 4}) // true

// Check if slice is reverse sorted
isReverse := pkg.IsSortedSliceDesc([]int{4, 3, 2, 1}) // true

// Validate number range
valid := pkg.IsValidRange(50, 1, 100) // true

// Check if slice is empty
empty := pkg.IsEmpty([]int{}) // true

// Check for duplicates
hasDups := pkg.HasDuplicates([]int{1, 2, 2, 3}) // true

// Count duplicates
count := pkg.CountDuplicates([]int{1, 2, 2, 3, 3}) // 2

// Find min/max values
min, max, valid := pkg.FindMinMax([]int{3, 1, 4, 1, 5}) // 1, 5, true
```

### 📊 **Performance Module** (`performance.go`)

Provides utilities for performance analysis and benchmarking.

**Key Types:**
- `PerformanceAnalysis` - Performance metrics
- `BenchmarkResult` - Individual benchmark result
- `ScalingAnalysis` - Performance scaling analysis
- `BenchmarkSummary` - Collection of benchmark results

**Key Functions:**
```go
// Calculate performance analysis
analysis := pkg.CalculateAnalysis(count, duration)

// Calculate scaling between benchmarks
scaling := pkg.CalculateScaling(results)

// Print basic performance info
pkg.PrintPerformanceInfo(count, duration, analysis)

// Print detailed performance analysis
pkg.PrintDetailedPerformance(count, duration, analysis, isSorted)

// Print comprehensive benchmark summary
pkg.PrintBenchmarkSummary(summary)
```

## 🚀 Usage Examples

### Basic Input and Validation
```go
package main

import "github.com/JoaoVitor615/algorithms-in-go/pkg"

func main() {
    reader := pkg.NewInputReader()
    
    // Get validated number input
    count := reader.ReadIntOrDefault("Enter count (1-1000): ", 1, 1000)
    if count == -1 {
        println("Invalid input!")
        return
    }
    
    // Generate and validate data
    gen := pkg.NewRandomGenerator()
    numbers := gen.GenerateIntSliceDefault(count)
    
    if pkg.HasDuplicates(numbers) {
        println("Generated data contains duplicates")
    }
    
    // Display formatted results
    pkg.PrintSubHeader("Generated Numbers")
    pkg.PrintSlicePartial(numbers, 10)
}
```

### Performance Analysis
```go
package main

import (
    "time"
    "github.com/JoaoVitor615/algorithms-in-go/pkg"
)

func main() {
    // Generate test data
    gen := pkg.NewRandomGenerator()
    numbers := gen.GenerateIntSliceDefault(1000)
    
    // Measure sorting performance
    start := time.Now()
    // ... perform sorting ...
    duration := time.Since(start)
    
    // Analyze performance
    analysis := pkg.CalculateAnalysis(len(numbers), duration)
    isSorted := pkg.IsSortedSlice(numbers)
    
    // Display results
    pkg.PrintDetailedPerformance(len(numbers), duration, analysis, isSorted)
}
```

## 🎯 Design Principles

1. **Generic and Reusable**: Functions are designed to work with common data types and patterns
2. **No Side Effects**: Most functions are pure and don't modify input data
3. **Clear Interfaces**: Simple, predictable function signatures
4. **Error Handling**: Graceful handling of edge cases and invalid input
5. **Performance Focused**: Efficient implementations suitable for benchmarking

## 🔄 Integration

This package is designed to be used throughout the algorithms-in-go project to:

- **Reduce Code Duplication**: Common operations are centralized
- **Improve Maintainability**: Changes to utilities affect all consumers
- **Ensure Consistency**: Uniform behavior across different algorithm implementations
- **Simplify Testing**: Utilities can be tested independently

---
