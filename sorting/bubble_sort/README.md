# 🫧 Bubble Sort

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Algorithm](https://img.shields.io/badge/Algorithm-Bubble%20Sort-orange?style=for-the-badge)
![Complexity](https://img.shields.io/badge/Time-O(n²)-red?style=for-the-badge)
![Space](https://img.shields.io/badge/Space-O(1)-green?style=for-the-badge)
![Stable](https://img.shields.io/badge/Stable-Yes-brightgreen?style=for-the-badge)

**A comprehensive implementation of the Bubble Sort algorithm in Go**

</div>

---

## 📋 Table of Contents

- [🔍 Overview](#-overview)
- [⚡ Algorithm Variants](#-algorithm-variants)
- [📊 Complexity Analysis](#-complexity-analysis)
- [🚀 Usage Examples](#-usage-examples)
- [🧪 Testing](#-testing)
- [🎯 When to Use](#-when-to-use)

---

## 🔍 Overview

Bubble Sort is one of the simplest sorting algorithms to understand and implement. It works by repeatedly stepping through the list, comparing adjacent elements and swapping them if they are in the wrong order. The pass through the list is repeated until the list is sorted.

### 🌟 Key Characteristics

- **Simple Implementation**: Easy to understand and code
- **Stable Algorithm**: Maintains relative order of equal elements
- **In-Place Sorting**: Requires only O(1) extra memory space
- **Adaptive**: Can be optimized to perform better on nearly sorted data
- **Educational Value**: Excellent for learning sorting concepts

### 🔄 How It Works

1. **Compare Adjacent Elements**: Start with the first two elements
2. **Swap if Necessary**: If they're in wrong order, swap them
3. **Move to Next Pair**: Continue with the next adjacent pair
4. **Complete Pass**: After each complete pass, the largest element "bubbles up" to its correct position
5. **Repeat**: Continue until no swaps are needed

---

## ⚡ Algorithm Variants

This implementation provides multiple variants of Bubble Sort:

### 1. **Basic Bubble Sort** (`BubbleSort`)
```go
func BubbleSort(arr []int) []int
```
- **Description**: Standard implementation that creates a copy of the input
- **Time Complexity**: O(n²) in all cases
- **Space Complexity**: O(n) for the copy
- **Use Case**: When you need to preserve the original array

### 2. **Optimized Bubble Sort** (`BubbleSortOptimized`)
```go
func BubbleSortOptimized(arr []int) []int
```
- **Description**: Enhanced version that stops early if array becomes sorted
- **Time Complexity**: O(n) best case, O(n²) worst case
- **Space Complexity**: O(n) for the copy
- **Use Case**: When input might be nearly sorted

### 3. **In-Place Bubble Sort** (`BubbleSortInPlace`)
```go
func BubbleSortInPlace(arr []int)
```
- **Description**: Sorts the array in-place without creating a copy
- **Time Complexity**: O(n²) in all cases
- **Space Complexity**: O(1)
- **Use Case**: When memory is limited

### 4. **Optimized In-Place** (`BubbleSortInPlaceOptimized`)
```go
func BubbleSortInPlaceOptimized(arr []int)
```
- **Description**: Best of both worlds - in-place and optimized
- **Time Complexity**: O(n) best case, O(n²) worst case
- **Space Complexity**: O(1)
- **Use Case**: Most practical variant for small datasets

### 5. **With Callback** (`BubbleSortWithCallback`)
```go
func BubbleSortWithCallback(arr []int, callback func([]int, int, int)) []int
```
- **Description**: Educational variant that calls a function after each swap
- **Use Case**: Visualization, debugging, or educational purposes

---

## 📊 Complexity Analysis

| Variant | Best Case | Average Case | Worst Case | Space | Stable |
|---------|-----------|--------------|------------|-------|--------|
| **Basic** | O(n²) | O(n²) | O(n²) | O(n) | ✅ Yes |
| **Optimized** | O(n) | O(n²) | O(n²) | O(n) | ✅ Yes |
| **In-Place** | O(n²) | O(n²) | O(n²) | O(1) | ✅ Yes |
| **Optimized In-Place** | O(n) | O(n²) | O(n²) | O(1) | ✅ Yes |

### 🎯 Performance Scenarios

- **Best Case**: Already sorted array (optimized versions only)
- **Average Case**: Random order array
- **Worst Case**: Reverse sorted array

---

## 🚀 Usage Examples

### 📝 **Basic Usage**

```go
package main

import (
    "fmt"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/bubble_sort"
)

func main() {
    // Basic Bubble Sort
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sorted := bubble_sort.BubbleSort(arr)
    fmt.Println("Original:", arr)    // [64 34 25 12 22 11 90]
    fmt.Println("Sorted:", sorted)   // [11 12 22 25 34 64 90]
}
```

### ⚡ **Optimized Version**

```go
func main() {
    // Nearly sorted array - optimized version performs much better
    arr := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}
    sorted := bubble_sort.BubbleSortOptimized(arr)
    fmt.Println("Sorted:", sorted)   // [1 2 3 4 5 6 7 8 9]
}
```

### 💾 **In-Place Sorting**

```go
func main() {
    // In-place sorting - modifies original array
    arr := []int{5, 2, 8, 1, 9}
    bubble_sort.BubbleSortInPlace(arr)
    fmt.Println("Sorted in-place:", arr)  // [1 2 5 8 9]
}
```

### 🎨 **With Visualization Callback**

```go
func main() {
    arr := []int{4, 2, 7, 1, 3}
    
    // Callback function to visualize swaps
    callback := func(array []int, i, j int) {
        fmt.Printf("Swapped positions %d and %d: %v\n", i, j, array)
    }
    
    sorted := bubble_sort.BubbleSortWithCallback(arr, callback)
    fmt.Println("Final result:", sorted)
}
```

### 📊 **Performance Comparison**

```go
func main() {
    import "time"
    
    // Large dataset
    data := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        data[i] = 1000 - i  // Reverse sorted (worst case)
    }
    
    // Test basic version
    start := time.Now()
    bubble_sort.BubbleSort(data)
    fmt.Printf("Basic version: %v\n", time.Since(start))
    
    // Test optimized version
    start = time.Now()
    bubble_sort.BubbleSortOptimized(data)
    fmt.Printf("Optimized version: %v\n", time.Since(start))
}
```

---

## 🧪 Testing

### 🎯 **Running Tests**

```bash
# Run all tests
go test ./sorting/bubble_sort

# Run tests with coverage
go test -cover ./sorting/bubble_sort

# Run benchmarks
go test -bench=. ./sorting/bubble_sort

# Run specific test
go test -run TestBubbleSortOptimized ./sorting/bubble_sort
```

### 📊 **Test Coverage**

The implementation includes comprehensive tests covering:

- ✅ **Edge Cases**: Empty arrays, single elements
- ✅ **Data Patterns**: Sorted, reverse sorted, random, duplicates
- ✅ **Algorithm Variants**: All 5 variants thoroughly tested
- ✅ **Stability**: Verification that equal elements maintain order
- ✅ **Performance**: Benchmarks for different input sizes
- ✅ **Memory Safety**: Original array preservation tests

### 🏃 **Benchmark Results**

```bash
BenchmarkBubbleSort/size_100-8         	    5000	    245670 ns/op
BenchmarkBubbleSort/size_1000-8        	      50	  24567000 ns/op
BenchmarkBubbleSortOptimized/sorted_size_100-8  	 1000000	      1234 ns/op
BenchmarkBubbleSortOptimized/reverse_size_100-8 	    5000	    245670 ns/op
```

---

## 🎯 When to Use

### ✅ **Good For:**
- **Educational Purposes**: Learning sorting algorithms
- **Small Datasets**: Arrays with < 50 elements
- **Nearly Sorted Data**: When using optimized version
- **Stable Sorting**: When relative order of equal elements matters
- **Memory Constraints**: When O(1) space complexity is required
- **Simple Implementation**: When code simplicity is prioritized

### ❌ **Avoid When:**
- **Large Datasets**: Performance degrades significantly (O(n²))
- **Performance Critical**: Other algorithms like Quick Sort or Merge Sort are faster
- **Production Systems**: Generally not suitable for production sorting needs

### 🔄 **Alternatives to Consider:**
- **Quick Sort**: Better average performance O(n log n)
- **Merge Sort**: Guaranteed O(n log n) and stable
- **Insertion Sort**: Better for small arrays and nearly sorted data
- **Heap Sort**: Guaranteed O(n log n) with O(1) space

---

## 🎓 Educational Value

Bubble Sort is excellent for learning because it demonstrates:

1. **Basic Sorting Concepts**: Comparison-based sorting
2. **Algorithm Optimization**: Early termination techniques
3. **Stability**: How to maintain relative order
4. **In-Place Operations**: Memory-efficient algorithms
5. **Time Complexity**: Understanding O(n²) behavior

---

<div align="center">

**Part of the [Sorting Algorithms](../README.md) collection**

Made with ❤️ and lots of ☕ by [JoaoVitor615](https://github.com/JoaoVitor615)

</div>

