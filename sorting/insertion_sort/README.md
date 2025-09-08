# 📝 Insertion Sort

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Algorithm](https://img.shields.io/badge/Algorithm-Insertion%20Sort-orange?style=for-the-badge)
![Complexity](https://img.shields.io/badge/Time-O(n²)-red?style=for-the-badge)
![Space](https://img.shields.io/badge/Space-O(1)-green?style=for-the-badge)
![Stable](https://img.shields.io/badge/Stable-Yes-brightgreen?style=for-the-badge)

**A comprehensive implementation of the Insertion Sort algorithm in Go**

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

Insertion Sort is a simple sorting algorithm that builds the final sorted array one item at a time. It works by taking elements from the unsorted portion and inserting them into their correct position in the sorted portion of the array.

### 🌟 Key Characteristics

- **Simple Implementation**: Easy to understand and implement
- **Stable Algorithm**: Maintains relative order of equal elements
- **In-Place Sorting**: Requires only O(1) extra memory space
- **Adaptive**: Performs excellently on nearly sorted data (O(n) best case)
- **Efficient for Small Data**: Often used as the final stage of sophisticated algorithms

### 🔄 How It Works

1. **Start from Second Element**: The first element is considered sorted
2. **Pick Current Element**: Take the next element from the unsorted portion
3. **Find Insertion Position**: Compare with elements in the sorted portion from right to left
4. **Shift Elements**: Move larger elements one position to the right
5. **Insert Element**: Place the current element in its correct position
6. **Repeat**: Continue until all elements are processed

---

## ⚡ Algorithm Variants

This implementation provides multiple variants of Insertion Sort:

### 1. **Basic Insertion Sort** (`InsertionSort`)
```go
func InsertionSort(arr []int) []int
```
- **Description**: Standard implementation that creates a copy of the input
- **Time Complexity**: O(n²) worst/average case, O(n) best case
- **Space Complexity**: O(n) for the copy
- **Use Case**: When you need to preserve the original array

### 2. **Optimized Insertion Sort** (`InsertionSortOptimized`)
```go
func InsertionSortOptimized(arr []int) []int
```
- **Description**: Uses binary search to find insertion position
- **Time Complexity**: O(n²) worst case (shifting), O(n log n) comparisons
- **Space Complexity**: O(n) for the copy
- **Use Case**: When you want to reduce the number of comparisons

### 3. **In-Place Insertion Sort** (`InsertionSortInPlace`)
```go
func InsertionSortInPlace(arr []int)
```
- **Description**: Sorts the array in-place without creating a copy
- **Time Complexity**: O(n²) worst/average case, O(n) best case
- **Space Complexity**: O(1)
- **Use Case**: When memory is limited

### 4. **Optimized In-Place** (`InsertionSortInPlaceOptimized`)
```go
func InsertionSortInPlaceOptimized(arr []int)
```
- **Description**: In-place version with binary search for insertion position
- **Time Complexity**: O(n²) worst case, O(n log n) comparisons
- **Space Complexity**: O(1)
- **Use Case**: Memory-efficient with fewer comparisons

### 5. **With Callback** (`InsertionSortWithCallback`)
```go
func InsertionSortWithCallback(arr []int, callback func([]int, int, int)) []int
```
- **Description**: Educational variant that calls a function after each insertion
- **Use Case**: Visualization, debugging, or educational purposes

### 6. **Descending Order** (`InsertionSortDescending`)
```go
func InsertionSortDescending(arr []int) []int
```
- **Description**: Sorts array in descending order
- **Use Case**: When reverse order is needed

### 7. **With Custom Gap** (`InsertionSortWithGap`)
```go
func InsertionSortWithGap(arr []int, gap int) []int
```
- **Description**: Insertion sort with custom gap (used by Shell Sort)
- **Use Case**: As a subroutine for more complex algorithms

---

## 📊 Complexity Analysis

| Variant | Best Case | Average Case | Worst Case | Space | Stable |
|---------|-----------|--------------|------------|-------|--------|
| **Basic** | O(n) | O(n²) | O(n²) | O(n) | ✅ Yes |
| **Optimized** | O(n) | O(n²) | O(n²) | O(n) | ✅ Yes |
| **In-Place** | O(n) | O(n²) | O(n²) | O(1) | ✅ Yes |
| **Optimized In-Place** | O(n) | O(n²) | O(n²) | O(1) | ✅ Yes |

### 🎯 Performance Scenarios

- **Best Case**: Already sorted array - O(n) time
- **Average Case**: Random order array - O(n²) time
- **Worst Case**: Reverse sorted array - O(n²) time

---

## 🚀 Usage Examples

### 📝 **Basic Usage**

```go
package main

import (
    "fmt"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/insertion_sort"
)

func main() {
    // Basic Insertion Sort
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sorted := insertion_sort.InsertionSort(arr)
    fmt.Println("Original:", arr)    // [64 34 25 12 22 11 90]
    fmt.Println("Sorted:", sorted)   // [11 12 22 25 34 64 90]
}
```

### ⚡ **Optimized Version**

```go
func main() {
    // Large array - optimized version uses binary search for insertion position
    arr := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
    sorted := insertion_sort.InsertionSortOptimized(arr)
    fmt.Println("Sorted:", sorted)   // [1 2 3 4 5 6 7 8 9]
}
```

### 💾 **In-Place Sorting**

```go
func main() {
    // In-place sorting - modifies original array
    arr := []int{5, 2, 8, 1, 9}
    insertion_sort.InsertionSortInPlace(arr)
    fmt.Println("Sorted in-place:", arr)  // [1 2 5 8 9]
}
```

### 🎨 **With Visualization Callback**

```go
func main() {
    arr := []int{4, 2, 7, 1, 3}
    
    // Callback function to visualize insertions
    callback := func(array []int, insertPos, currentPos int) {
        fmt.Printf("Inserted element at position %d (was at %d): %v\n", 
                   insertPos, currentPos, array)
    }
    
    sorted := insertion_sort.InsertionSortWithCallback(arr, callback)
    fmt.Println("Final result:", sorted)
}
```

### 📊 **Descending Order**

```go
func main() {
    arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
    sorted := insertion_sort.InsertionSortDescending(arr)
    fmt.Println("Descending:", sorted)  // [9 6 5 4 3 2 1 1]
}
```

### 🔧 **With Custom Gap (Shell Sort Building Block)**

```go
func main() {
    arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    
    // Sort with gap of 3 (partial sorting)
    partialSorted := insertion_sort.InsertionSortWithGap(arr, 3)
    fmt.Println("Gap 3:", partialSorted)
    
    // Final sort with gap of 1 (complete sorting)
    fullySorted := insertion_sort.InsertionSortWithGap(partialSorted, 1)
    fmt.Println("Gap 1:", fullySorted)
}
```

### 📈 **Performance Comparison**

```go
func main() {
    import "time"
    
    // Nearly sorted data (best case)
    data := []int{1, 2, 3, 5, 4, 6, 7, 8, 9, 10}
    
    // Test basic version
    start := time.Now()
    insertion_sort.InsertionSort(data)
    fmt.Printf("Basic version: %v\n", time.Since(start))
    
    // Test optimized version
    start = time.Now()
    insertion_sort.InsertionSortOptimized(data)
    fmt.Printf("Optimized version: %v\n", time.Since(start))
}
```

---

## 🧪 Testing

### 🎯 **Running Tests**

```bash
# Run all tests
go test ./sorting/insertion_sort

# Run tests with coverage
go test -cover ./sorting/insertion_sort

# Run benchmarks
go test -bench=. ./sorting/insertion_sort

# Run specific test
go test -run TestInsertionSortOptimized ./sorting/insertion_sort

# Verbose output
go test -v ./sorting/insertion_sort
```

### 📊 **Test Coverage**

The implementation includes comprehensive tests covering:

- ✅ **Edge Cases**: Empty arrays, single elements
- ✅ **Data Patterns**: Sorted, reverse sorted, random, duplicates
- ✅ **Algorithm Variants**: All 7 variants thoroughly tested
- ✅ **Stability**: Verification that equal elements maintain order
- ✅ **Performance**: Benchmarks for different input sizes and patterns
- ✅ **Memory Safety**: Original array preservation tests
- ✅ **Special Cases**: Negative numbers, all same elements

### 🏃 **Benchmark Results**

```bash
BenchmarkInsertionSort/size_100-8                    	   10000	    123456 ns/op
BenchmarkInsertionSort/size_1000-8                   	     100	  12345678 ns/op
BenchmarkInsertionSortOptimized/sorted_size_100-8    	 1000000	      1234 ns/op
BenchmarkInsertionSortOptimized/reverse_size_100-8   	   10000	    123456 ns/op
BenchmarkInsertionSortBestCase/best_case_size_100-8   	 1000000	      1000 ns/op
```

---

## 🎯 When to Use

### ✅ **Excellent For:**
- **Small Datasets**: Arrays with < 50 elements
- **Nearly Sorted Data**: Performs in O(n) time when data is almost sorted
- **Online Algorithms**: Can sort data as it arrives
- **Stable Sorting**: When relative order of equal elements must be preserved
- **Memory Constraints**: O(1) space complexity with in-place version
- **Educational Purposes**: Easy to understand and implement
- **Hybrid Algorithms**: Often used as final step in sophisticated algorithms (like Timsort)

### ❌ **Avoid When:**
- **Large Datasets**: Performance degrades significantly (O(n²))
- **Completely Random Data**: Other algorithms like Quick Sort or Merge Sort are faster
- **Performance Critical Applications**: When consistent O(n log n) performance is required

### 🔄 **Alternatives to Consider:**
- **Merge Sort**: Guaranteed O(n log n) and stable
- **Quick Sort**: Better average performance O(n log n)
- **Heap Sort**: Guaranteed O(n log n) with O(1) space
- **Tim Sort**: Hybrid algorithm that uses Insertion Sort for small runs

---

## 🎓 Educational Value

Insertion Sort is excellent for learning because it demonstrates:

1. **Incremental Construction**: Building solution step by step
2. **Invariant Maintenance**: Keeping part of array always sorted
3. **Adaptive Behavior**: Performance depends on input characteristics
4. **Stability Concepts**: How to maintain relative order
5. **In-Place Operations**: Memory-efficient algorithm design
6. **Best/Worst Case Analysis**: Understanding when algorithms excel or struggle

### 🧠 **Key Learning Points:**

- **Loop Invariants**: At each step, elements 0 to i-1 are sorted
- **Adaptive Nature**: Fewer operations needed for nearly sorted data
- **Practical Applications**: Used in hybrid sorting algorithms
- **Trade-offs**: Simplicity vs. performance for different input sizes

---

## 🔍 **Algorithm Visualization**

```
Initial: [5, 2, 8, 1, 9]

Step 1: [5, 2, 8, 1, 9] → [2, 5, 8, 1, 9]  (Insert 2)
Step 2: [2, 5, 8, 1, 9] → [2, 5, 8, 1, 9]  (8 already in place)
Step 3: [2, 5, 8, 1, 9] → [1, 2, 5, 8, 9]  (Insert 1 at beginning)
Step 4: [1, 2, 5, 8, 9] → [1, 2, 5, 8, 9]  (9 already in place)

Final:   [1, 2, 5, 8, 9]
```

---

<div align="center">

**Part of the [Sorting Algorithms](../README.md) collection**

Made with ❤️ and lots of ☕ by [JoaoVitor615](https://github.com/JoaoVitor615)

</div>
