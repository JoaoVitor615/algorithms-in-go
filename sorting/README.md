# 🔀 Sorting Algorithms

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Algorithms](https://img.shields.io/badge/Algorithms-Sorting-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)

**A comprehensive collection of sorting algorithms implemented in Go**

</div>

---

## 📋 Table of Contents

- [🔍 Overview](#-overview)
- [🏗️ Architecture](#️-architecture)
- [🔗 Algorithms Implemented](#-algorithms-implemented)
- [🚀 Usage](#-usage)
- [🧪 Testing](#-testing)
- [🤝 Contributing](#-contributing)

---

## 🔍 Overview

This module contains a comprehensive collection of sorting algorithms implemented in Go. The architecture is designed to be modular, extensible, and educational, making it easy to add new sorting algorithms while maintaining a consistent interface.

### 🌟 Features

- **Modular Architecture**: Each algorithm is self-contained in its own package
- **Common Interface**: Unified terminal and use case layers for all sorting algorithms
- **Comprehensive Testing**: Each algorithm includes extensive test coverage
- **Performance Analysis**: Built-in benchmarking and complexity analysis
- **Interactive Demo**: Terminal-based interface for hands-on experimentation

---

## 🏗️ Architecture

```
sorting/
├── terminal.go              # Common terminal interface for all sorting algorithms
├── use_cases.go            # Common business logic and use cases
├── README.md               # This documentation
├── merge_sort/             # Merge Sort implementation
│   ├── mergesort.go        # Core algorithm
│   └── mergesort_test.go   # Algorithm tests
├── quick_sort/             # Quick Sort implementation
├── bubble_sort/            # Bubble Sort implementation  
├── heap_sort/              # Heap Sort (Coming Soon)
└── insertion_sort/         # Insertion Sort (Coming Soon)
```

### 🎯 Design Principles

- **Separation of Concerns**: Algorithm logic separated from UI and business logic
- **Single Responsibility**: Each file has a clear, focused purpose
- **Extensibility**: Easy to add new algorithms without modifying existing code
- **Consistency**: All algorithms follow the same patterns and interfaces

---

## 🔗 Algorithms Implemented

| Algorithm | Time Complexity | Space Complexity | Stable | Status |
|-----------|----------------|------------------|--------|---------|
| **Merge Sort** | O(n log n) | O(log n) | ✅ Yes | ✅ Implemented |
| **Quick Sort** | O(n log n) avg, O(n²) worst | O(log n) | ❌ No | ✅ Implemented |
| **Bubble Sort** | O(n²) avg, O(n) best | O(1) | ✅ Yes | ✅ Implemented |
| **Heap Sort** | O(n log n) | O(1) | ❌ No | 🔄 Coming Soon |
| **Insertion Sort** | O(n²) | O(1) | ✅ Yes | 🔄 Coming Soon |

### 📊 Algorithm Details

#### ✅ **Merge Sort**
- **Type**: Divide and conquer
- **Data Structure**: Linked lists
- **Best for**: Large datasets, stable sorting requirements, external sorting
- **Implementation**: Linked list based for optimal memory usage
- **Features**: Performance analysis, benchmarking, visualization

#### ✅ **Quick Sort**
- **Type**: Divide and conquer with partitioning
- **Data Structure**: Arrays
- **Best for**: General purpose sorting, average case performance
- **Implementation**: In-place partitioning with configurable pivot strategies
- **Features**: Multiple pivot strategies, performance analysis, benchmarking

#### ✅ **Bubble Sort**
- **Type**: Comparison-based with adjacent swapping
- **Data Structure**: Arrays
- **Best for**: Educational purposes, small datasets, nearly sorted data
- **Implementation**: Multiple variants including optimized and in-place versions
- **Features**: Early termination optimization, visualization callbacks, comprehensive testing

---

## 🚀 Usage

### 📝 **Basic Usage**

```go
package main

import (
    "fmt"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/quick_sort"
    "github.com/JoaoVitor615/algorithms-in-go/sorting/bubble_sort"
)

func main() {
    // Merge Sort (Linked Lists)
    head := &merge_sort.Node{Value: 64}
    head.Next = &merge_sort.Node{Value: 34}
    head.Next.Next = &merge_sort.Node{Value: 25}
    
    sortedHead := merge_sort.MergeSort(head)
    
    // Quick Sort (Arrays)
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sortedArr := quick_sort.QuickSort(arr)
    fmt.Println(sortedArr) // [11 12 22 25 34 64 90]
    
    // Bubble Sort (Arrays)
    arr2 := []int{64, 34, 25, 12, 22, 11, 90}
    sortedArr2 := bubble_sort.BubbleSortOptimized(arr2)
    fmt.Println(sortedArr2) // [11 12 22 25 34 64 90]
}
```

### 🎮 **Interactive Interface**

```go
package main

import "github.com/JoaoVitor615/algorithms-in-go/sorting"

func main() {
    // Launch interactive sorting interface
    sorting.RunSortingInterface()
}
```

### 📊 **Advanced Testing Options**

The sorting interface provides 8 different testing modes:

1. **Manual Input** - Enter numbers manually for educational purposes
2. **Custom Random** - Generate 1 to 1,000,000 random numbers
3. **Benchmark 500** - Standard small dataset benchmark
4. **Benchmark 1,000** - Medium dataset benchmark
5. **Benchmark 5,000** - Large dataset benchmark
6. **Benchmark 10,000** - Extra large dataset benchmark
7. **All Benchmarks** - Comprehensive performance analysis
8. **Back to Menu** - Return to algorithm selection

### 🔧 **Example Session**

```
[   Sorting Algorithms - Advanced Testing   ]
Choose a sorting algorithm:

1. Merge Sort
2. Quick Sort
3. Bubble Sort
4. Heap Sort (Coming Soon)
5. Insertion Sort (Coming Soon)
6. Back to main menu

Enter your choice (1-6): 3

[   Bubble Sort - Advanced Testing   ]
Choose a testing option:

1. Manual input (enter numbers manually)
2. Custom random list (specify quantity)
3. Benchmark 500 random numbers
4. Benchmark 1,000 random numbers
5. Benchmark 5,000 random numbers
6. Benchmark 10,000 random numbers
7. Run all benchmarks
8. Back to sorting menu

Enter your choice (1-8): 7

=== Bubble Sort - Running All Benchmarks ===
This will test sorting performance with different input sizes...

[1/4] Generating 500 numbers... Sorting... Done in 24.1µs ✅
[2/4] Generating 1,000 numbers... Sorting... Done in 49.5µs ✅
[3/4] Generating 5,000 numbers... Sorting... Done in 304.7µs ✅
[4/4] Generating 10,000 numbers... Sorting... Done in 637.2µs ✅

============================================================
               BENCHMARK SUMMARY
============================================================
Count        Time            μs/number       Status    
------------------------------------------------------------
500          24.1µs          0.05            ✅         
1,000        49.5µs          0.05            ✅         
5,000        304.7µs         0.06            ✅         
10,000       637.2µs         0.06            ✅         
============================================================

📊 SCALING ANALYSIS:
   500 → 1,000: 2.00x size, 2.06x time
   1,000 → 5,000: 5.00x size, 6.16x time
   5,000 → 10,000: 2.00x size, 2.09x time

💡 Note: Performance depends on the algorithm complexity.
   For O(n log n) algorithms: Expected time ratio for 2x size: ~2.2x time
```

---

## 🧪 Testing

### 🎯 **Running Tests**

```bash
# Test all sorting algorithms
go test ./sorting/...

# Test specific algorithm
go test ./sorting/merge_sort

# Test with coverage
go test -cover ./sorting/...

# Run benchmarks
go test -bench=. ./sorting/merge_sort
```

### 📊 **Test Coverage**

| Algorithm | Coverage | Test Cases | Status |
|-----------|----------|------------|---------|
| Merge Sort | 43.3% | 7 comprehensive cases | ✅ Complete |
| Quick Sort | 85%+ | 10+ comprehensive cases | ✅ Complete |
| Bubble Sort | 95%+ | 15+ comprehensive cases | ✅ Complete |
| Heap Sort | - | - | 🔄 Coming Soon |

---

## 🤝 Contributing

### 📋 **Adding New Algorithms**

1. **Create Algorithm Directory**:
   ```bash
   mkdir sorting/your_algorithm/
   ```

2. **Implement Core Algorithm**:
   ```go
   // sorting/your_algorithm/algorithm.go
   package your_algorithm
   
   import "github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"
   
   func YourSort(head *merge_sort.Node) *merge_sort.Node {
       // Your implementation here
   }
   ```

3. **Add Tests**:
   ```go
   // sorting/your_algorithm/algorithm_test.go
   package your_algorithm
   
   func TestYourSort(t *testing.T) {
       // Your tests here
   }
   ```

4. **Update Common Files**:
   - Add case in `sorting/use_cases.go` → `executeSort()` function
   - Add menu option in `sorting/terminal.go` → `showSortingMenu()` function

### 🔧 **Algorithm Template**

```go
package your_algorithm

import "github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"

// YourSort implements [brief description]
// Time Complexity: O(?)
// Space Complexity: O(?)
// Stable: Yes/No
func YourSort(head *merge_sort.Node) *merge_sort.Node {
    // Base case
    if head == nil || head.Next == nil {
        return head
    }
    
    // Your sorting logic here
    
    return head // Return sorted list
}
```

---

<div align="center">

**Part of the [Algorithms in Go](../README.md) collection**

Made with ❤️ and lots of ☕ by [JoaoVitor615](https://github.com/JoaoVitor615)

</div>
