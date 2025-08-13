# 🚀 Algorithms in Go

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/github/license/JoaoVitor615/algorithms-in-go?style=for-the-badge)
![GitHub stars](https://img.shields.io/github/stars/JoaoVitor615/algorithms-in-go?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/JoaoVitor615/algorithms-in-go?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/JoaoVitor615/algorithms-in-go?style=for-the-badge)

**A comprehensive collection of algorithms and data structures implemented in Go**

[🔗 Explore the Code](#-algorithms-implemented) • [📖 Documentation](#-documentation) • [🛠️ Installation](#️-installation) • [🤝 Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [📝 About](#-about)
- [🎯 Objectives](#-objectives)
- [🔗 Algorithms Implemented](#-algorithms-implemented)
- [🛠️ Installation](#️-installation)
- [🚀 Usage](#-usage)
- [🧪 Testing](#-testing)
- [📈 Performance Analysis](#-performance-analysis)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👨‍💻 Author](#-author)

---

## 📝 About

**Algorithms in Go** is an educational repository designed to help developers understand and implement fundamental algorithms and data structures using the Go programming language. This project serves as both a learning resource and a reference implementation for computer science concepts.

### 🌟 Why This Repository?

- **Educational Focus**: Clear, well-documented implementations perfect for learning
- **Production Ready**: Efficient algorithms with proper error handling
- **Test Coverage**: Comprehensive test suites ensuring reliability
- **Interactive Examples**: Terminal-based interfaces for hands-on experimentation
- **Community Driven**: Open for contributions and improvements

---

## 🎯 Objectives

This repository aims to:

- 📚 **Educate**: Provide clear implementations of fundamental algorithms
- 🔬 **Demonstrate**: Show Go language best practices and idiomatic code
- 🚀 **Inspire**: Encourage algorithmic thinking and problem-solving skills
- 🤝 **Build Community**: Foster collaboration among Go developers
- 📊 **Analyze**: Compare algorithm performance and complexity

---

## 🔗 Algorithms Implemented

<details>
<summary><strong>🔀 Sorting Algorithms</strong></summary>

| Algorithm | Time Complexity | Space Complexity | Stable | Status |
|-----------|----------------|------------------|--------|---------|
| **Merge Sort** | O(n log n) | O(n) | ✅ | ✅ Implemented |
| Quick Sort | O(n log n) avg, O(n²) worst | O(log n) | ❌ | 🔄 Coming Soon |
| Heap Sort | O(n log n) | O(1) | ❌ | 🔄 Coming Soon |
| Bubble Sort | O(n²) | O(1) | ✅ | 🔄 Coming Soon |
| Insertion Sort | O(n²) | O(1) | ✅ | 🔄 Coming Soon |

</details>

<details>
<summary><strong>🔍 Search Algorithms</strong></summary>

| Algorithm | Time Complexity | Space Complexity | Prerequisites | Status |
|-----------|----------------|------------------|---------------|---------|
| Binary Search | O(log n) | O(1) | Sorted Array | 🔄 Coming Soon |
| Linear Search | O(n) | O(1) | None | 🔄 Coming Soon |
| Depth-First Search (DFS) | O(V + E) | O(V) | Graph/Tree | 🔄 Coming Soon |
| Breadth-First Search (BFS) | O(V + E) | O(V) | Graph/Tree | 🔄 Coming Soon |

</details>

<details>
<summary><strong>🏗️ Data Structures</strong></summary>

| Structure | Average Access | Average Insert | Average Delete | Status |
|-----------|---------------|---------------|---------------|---------|
| Linked List | O(n) | O(1) | O(1) | ✅ Implemented |
| Stack | O(n) | O(1) | O(1) | 🔄 Coming Soon |
| Queue | O(n) | O(1) | O(1) | 🔄 Coming Soon |
| Binary Tree | O(log n) | O(log n) | O(log n) | 🔄 Coming Soon |
| Hash Table | O(1) | O(1) | O(1) | 🔄 Coming Soon |

</details>

<details>
<summary><strong>🔢 Dynamic Programming</strong></summary>

| Problem | Approach | Time Complexity | Status |
|---------|----------|----------------|---------|
| Fibonacci | Memoization | O(n) | 🔄 Coming Soon |
| Longest Common Subsequence | Bottom-up | O(mn) | 🔄 Coming Soon |
| Knapsack Problem | 2D DP | O(nW) | 🔄 Coming Soon |

</details>

---

## 🛠️ Installation

### Prerequisites

- **Go 1.24.3+** ([Download here](https://golang.org/dl/))
- **Git** for cloning the repository

### Quick Start

```bash
# Clone the repository
git clone https://github.com/JoaoVitor615/algorithms-in-go.git

# Navigate to the project directory
cd algorithms-in-go

# Initialize Go modules (if needed)
go mod tidy

# Run the interactive terminal
go run main.go
```

### Alternative: Direct Package Usage

```bash
# Install as a Go module
go get github.com/JoaoVitor615/algorithms-in-go
```

---

## 🚀 Usage

### Interactive Terminal Mode

The repository includes an interactive terminal interface for hands-on experimentation:

```bash
go run main.go
```

**Sample Output:**
```
Welcome to the Algorithms-in-Go Terminal! 🚀
Please choose an algorithm to execute:
1. Merge Sort
2. (Other algorithms will go here)

Enter your choice: 1
```

### Merge Sort Example

```go
package main

import (
    "fmt"
    mergesort "github.com/JoaoVitor615/algorithms-in-go/MergeSort"
)

func main() {
    // Create a linked list: 64 -> 34 -> 25 -> 12 -> nil
    head := &mergesort.Node{Value: 64}
    head.Next = &mergesort.Node{Value: 34}
    head.Next.Next = &mergesort.Node{Value: 25}
    head.Next.Next.Next = &mergesort.Node{Value: 12}

    fmt.Println("Original List:")
    printList(head) // Output: 64 -> 34 -> 25 -> 12 -> nil

    // Sort the list
    sortedHead := mergesort.MergeSort(head)

    fmt.Println("Sorted List:")
    printList(sortedHead) // Output: 12 -> 25 -> 34 -> 64 -> nil
}

func printList(node *mergesort.Node) {
    for node != nil {
        fmt.Printf("%d -> ", node.Value)
        node = node.Next
    }
    fmt.Println("nil")
}
```

### Advanced Testing Interface

The project includes an advanced testing interface with multiple options:

```go
package main

import mergesort "github.com/JoaoVitor615/algorithms-in-go/MergeSort"

func main() {
    // Launch advanced testing interface
    mergesort.RunAdvancedInterface()
}
```

**Available Testing Options:**
1. **Manual Input** - Enter numbers manually
2. **Custom Random** - Generate 1 to 1,000,000 random numbers
3. **Benchmark 500** - Test with 500 random numbers
4. **Benchmark 1,000** - Test with 1,000 random numbers  
5. **Benchmark 5,000** - Test with 5,000 random numbers
6. **Benchmark 10,000** - Test with 10,000 random numbers
7. **Run All Benchmarks** - Complete performance analysis
8. **Back to Main Menu** - Return to algorithm selection

### Using Individual Algorithms

```go
import "github.com/JoaoVitor615/algorithms-in-go/MergeSort"

// Use MergeSort directly
sortedList := mergesort.MergeSort(yourLinkedList)

// Or use the simple terminal interface (legacy)
mergesort.RunTerminal()
```

---

## 🧪 Testing

Run comprehensive tests to ensure algorithm correctness:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage report
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

# Run specific package tests
go test ./MergeSort -v
```

**Current Test Coverage:**
- **MergeSort**: 43.3% coverage with comprehensive edge case testing

**Test Cases Include:**
- ✅ Empty lists
- ✅ Single element lists
- ✅ Already sorted lists
- ✅ Reverse sorted lists
- ✅ Lists with duplicate elements
- ✅ Lists with even/odd number of elements

---

## 📈 Performance Analysis

### Merge Sort Performance

| Input Size | Time Complexity | Space Complexity | Performance |
|------------|----------------|------------------|-------------|
| n = 100 | O(n log n) | O(n) | ~0.001s |
| n = 1,000 | O(n log n) | O(n) | ~0.01s |
| n = 10,000 | O(n log n) | O(n) | ~0.1s |
| n = 100,000 | O(n log n) | O(n) | ~1s |

**Benchmarking:**
```bash
# Run performance benchmarks
go test -bench=. ./MergeSort
```

---

## 🤝 Contributing

We welcome contributions from the community! Here's how you can help:

### 🔧 Areas for Contribution

- **New Algorithms**: Implement additional sorting, searching, or graph algorithms
- **Optimization**: Improve existing algorithm performance
- **Documentation**: Enhance code comments and README sections
- **Testing**: Add more comprehensive test cases
- **Visualization**: Create graphical representations of algorithms

### 📋 Contribution Guidelines

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-algorithm`)
3. **Follow** Go conventions and add comprehensive tests
4. **Document** your code with clear comments
5. **Test** your implementation (`go test ./...`)
6. **Commit** your changes (`git commit -m 'Add amazing algorithm'`)
7. **Push** to the branch (`git push origin feature/amazing-algorithm`)
8. **Open** a Pull Request

### 💡 Algorithm Implementation Template

```go
package youralgorithm

// YourAlgorithm implements [brief description]
// Time Complexity: O(?)
// Space Complexity: O(?)
func YourAlgorithm(input []int) []int {
    // Implementation here
}

// Helper functions
func helper() {
    // Helper implementation
}
```

---

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2025 JoaoVitor615

Permission is hereby granted, free of charge, to any person obtaining a copy...
```

---

## 👨‍💻 Author

<div align="center">

### **João Vitor** 


**Passionate Go Developer  | Open Source Contributor**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/jvboaventura/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/JoaoVitor615)

</div>

### 🎯 Mission

> "Sharing knowledge about algorithms and data structures to help developers grow and build better software solutions."

### 🚀 Let's Connect!

I'm always excited to discuss algorithms, Go programming, and software engineering. Feel free to:

- 💬 **Ask questions** about implementations
- 🤝 **Collaborate** on new algorithms
- 📈 **Share** performance optimization ideas
- 🌟 **Star** this repository if it helps you!

---

<div align="center">

### ⭐ If this project helped you, please give it a star!

**Made with ❤️ and lots of ☕ by [JoaoVitor615](https://github.com/JoaoVitor615)**

---

<img src="https://github.com/betandr/gophers/blob/master/Gopher.png" alt="Go Gopher" height="300">

</div>
