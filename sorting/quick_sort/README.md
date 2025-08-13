# Quick Sort Implementation in Go

## 📖 Overview

This directory contains a comprehensive implementation of the **Quick Sort** algorithm in Go, featuring array-based sorting with multiple pivot strategies and advanced performance analysis.

---

## 🎯 Quick Sort Algorithm

**Quick Sort** is a highly efficient divide-and-conquer sorting algorithm that works by selecting a 'pivot' element and partitioning the array around it.

### 📊 Algorithm Characteristics

| Property | Value |
|----------|-------|
| **Type** | Divide and Conquer |
| **Data Structure** | Arrays |
| **Time Complexity (Average)** | O(n log n) |
| **Time Complexity (Worst)** | O(n²) |
| **Space Complexity** | O(log n) |
| **Stability** | ❌ No |
| **In-Place** | ✅ Yes |

### 🔄 How It Works

1. **Choose Pivot**: Select an element as the pivot (various strategies available)
2. **Partition**: Rearrange array so elements smaller than pivot come before it, and elements greater come after
3. **Recursively Sort**: Apply Quick Sort to the sub-arrays on both sides of the pivot
4. **Combine**: The array is sorted in-place, no explicit combine step needed

---

## 🛠️ Implementation Features

### 🎯 **Multiple Algorithms**

- **`QuickSort(arr []int) []int`**: Standard Quick Sort returning a new sorted array
- **`QuickSortInPlace(arr []int)`**: In-place sorting that modifies the original array
- **`QuickSortCustom(arr []int, strategy PivotStrategy) []int`**: Quick Sort with configurable pivot selection

### 🔄 **Pivot Strategies**

```go
type PivotStrategy int

const (
    LastElement   PivotStrategy = iota  // Choose last element (default)
    FirstElement                       // Choose first element
    MiddleElement                      // Choose middle element
    RandomElement                      // Choose random element
)
```

### ⚡ **Performance Optimizations**

- **Efficient Partitioning**: Lomuto partition scheme
- **Tail Recursion**: Optimized for better stack usage
- **Multiple Pivot Options**: Choose best strategy for your data

---

## 🚀 Usage

### 📝 **Basic Usage**

```go
package main

import (
    "fmt"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/quick_sort"
)

func main() {
    // Original array
    arr := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50}
    
    // Sort array (creates new array)
    sortedArr := quick_sort.QuickSort(arr)
    
    fmt.Println("Original:", arr)
    fmt.Println("Sorted:  ", sortedArr)
    // Output: [11 12 22 25 34 50 64 76 88 90]
}
```

### 🔧 **In-Place Sorting**

```go
func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    
    fmt.Println("Before:", arr)
    quick_sort.QuickSortInPlace(arr)
    fmt.Println("After: ", arr)
    // Output: [11 12 22 25 34 64 90]
}
```

### 🎯 **Custom Pivot Strategy**

```go
func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    
    // Different pivot strategies
    result1 := quick_sort.QuickSortCustom(arr, quick_sort.FirstElement)
    result2 := quick_sort.QuickSortCustom(arr, quick_sort.MiddleElement)
    result3 := quick_sort.QuickSortCustom(arr, quick_sort.RandomElement)
    
    fmt.Println("First Element Pivot: ", result1)
    fmt.Println("Middle Element Pivot:", result2)
    fmt.Println("Random Element Pivot:", result3)
}
```

---

## 🧪 Testing

### ✅ **Unit Tests**

Run the comprehensive test suite:

```bash
go test ./sorting/quick_sort
```

**Test Coverage:**
- Empty arrays
- Single element arrays
- Already sorted arrays
- Reverse sorted arrays
- Arrays with duplicates
- Arrays with negative numbers
- Large random arrays
- All pivot strategies

### 📊 **Benchmarks**

Run performance benchmarks:

```bash
go test -bench=. ./sorting/quick_sort
```

**Benchmark Features:**
- Multiple array sizes (100, 1,000, 10,000 elements)
- Worst-case scenarios (reverse sorted)
- Performance analysis and comparison

### 🎮 **Interactive Terminal Mode**

Test the algorithm interactively:

```bash
go run main.go
```

Navigate to: **Sorting Algorithms** → **Quick Sort**

#### Available Options:

1. **Manual Input**: Enter numbers manually to see step-by-step sorting
2. **Custom Random**: Generate custom list (1 to 1,000,000 numbers)
3. **Benchmark 500**: Test with 500 random numbers
4. **Benchmark 1,000**: Test with 1,000 random numbers  
5. **Benchmark 5,000**: Test with 5,000 random numbers
6. **Benchmark 10,000**: Test with 10,000 random numbers
7. **Run All Benchmarks**: Comprehensive performance analysis
8. **Back to Menu**: Return to sorting algorithm selection

---

## 📈 Performance Analysis

### ⏱️ **Time Complexity**

| Case | Complexity | Description |
|------|------------|-------------|
| **Best** | O(n log n) | Pivot divides array into equal halves |
| **Average** | O(n log n) | Typical random data performance |
| **Worst** | O(n²) | Already sorted or reverse sorted data |

### 💾 **Space Complexity**

- **Recursive Stack**: O(log n) average, O(n) worst case
- **In-Place**: No additional array space needed
- **Memory Efficient**: Only uses stack space for recursion

### 🎯 **When to Use Quick Sort**

**✅ Great for:**
- General purpose sorting
- Large datasets with random distribution
- When average O(n log n) performance is acceptable
- Memory-constrained environments (in-place sorting)
- When stability is not required

**❌ Consider alternatives for:**
- Nearly sorted data (poor performance)
- When worst-case O(n²) is unacceptable
- When stable sorting is required
- Small datasets (overhead may not be worth it)

---

## 📊 Example Output

### Interactive Demo

```
[   Quick Sort - Advanced Testing   ]
Choose a testing option:

1. Manual input (enter numbers manually)
2. Custom random list (specify quantity)
3. Benchmark 500 random numbers
4. Benchmark 1,000 random numbers
5. Benchmark 5,000 random numbers
6. Benchmark 10,000 random numbers
7. Run all benchmarks
8. Back to sorting menu

Enter your choice (1-8): 3

=== Quick Sort Benchmark: 500 Random Numbers ===
🎲 Generating 500 random numbers...
📝 List with 500 numbers generated successfully!
🔄 Starting sort...

✅ Sorting completed in: 71.167µs
📊 Performance: 500 numbers sorted
⚡ Average time per number: 0.14 μs
✅ Verification: List is correctly sorted!
📈 Theoretical O(n log n): 4483 operations
⏱️  Time per operation: 15.88 ns
```

---

## 🔗 Related Files

| File | Description |
|------|-------------|
| [`quicksort.go`](./quicksort.go) | Core Quick Sort implementation |
| [`quicksort_test.go`](./quicksort_test.go) | Comprehensive tests and benchmarks |
| [`../terminal.go`](../terminal.go) | Interactive terminal interface |
| [`../use_cases.go`](../use_cases.go) | Business logic and use cases |

---

## 🏗️ Architecture

This implementation follows Clean Architecture principles:

- **Algorithm Layer**: Pure sorting logic (`quicksort.go`)
- **Test Layer**: Comprehensive testing (`quicksort_test.go`)
- **Use Case Layer**: Business logic (`../use_cases.go`)
- **Interface Layer**: User interaction (`../terminal.go`)

---

## 🤝 Contributing

Feel free to contribute by:
- Adding new pivot strategies
- Optimizing performance
- Improving test coverage
- Enhancing documentation

---

*This Quick Sort implementation is part of the [Algorithms in Go](../../README.md) project.*
