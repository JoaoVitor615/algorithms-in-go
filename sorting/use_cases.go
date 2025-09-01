package sorting

import (
	"time"

	"github.com/JoaoVitor615/algorithms-in-go/pkg"
	"github.com/JoaoVitor615/algorithms-in-go/sorting/bubble_sort"
	"github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"
	"github.com/JoaoVitor615/algorithms-in-go/sorting/quick_sort"
)

// UseCase represents the business logic layer for sorting operations
type UseCase struct {
	generator *pkg.RandomGenerator
}

// NewUseCase creates a new UseCase instance
func NewUseCase() *UseCase {
	return &UseCase{
		generator: pkg.NewRandomGenerator(),
	}
}

// SortResult contains the result of a sorting operation
type SortResult struct {
	SortedList  *merge_sort.Node // For linked list algorithms
	SortedArray []int            // For array algorithms
	Duration    time.Duration
	Count       int
	IsSorted    bool
	Analysis    pkg.PerformanceAnalysis
	IsArray     bool // Flag to indicate if result is array or linked list
}

// Type aliases for convenience
type BenchmarkSummary = pkg.BenchmarkSummary
type BenchmarkResult = pkg.BenchmarkResult
type ScalingAnalysis = pkg.ScalingAnalysis

// ManualSort sorts a manually created list using the specified algorithm
func (uc *UseCase) ManualSort(algorithmName string, numbers []int) SortResult {
	if len(numbers) == 0 {
		return SortResult{
			SortedList:  nil,
			SortedArray: []int{},
			Duration:    0,
			Count:       0,
			IsSorted:    true,
			Analysis:    pkg.PerformanceAnalysis{},
			IsArray:     uc.algorithmUsesArray(algorithmName),
		}
	}

	count := len(numbers)
	startTime := time.Now()

	if uc.algorithmUsesArray(algorithmName) {
		// Array-based algorithms
		sortedArray := uc.executeSortArray(algorithmName, numbers)
		duration := time.Since(startTime)

		return SortResult{
			SortedList:  nil,
			SortedArray: sortedArray,
			Duration:    duration,
			Count:       count,
			IsSorted:    pkg.IsSortedSlice(sortedArray),
			Analysis:    pkg.CalculateAnalysis(count, duration),
			IsArray:     true,
		}
	} else {
		// Linked list algorithms
		head := uc.createListFromSlice(numbers)
		sortedList := uc.executeSortList(algorithmName, head)
		duration := time.Since(startTime)

		return SortResult{
			SortedList:  sortedList,
			SortedArray: nil,
			Duration:    duration,
			Count:       count,
			IsSorted:    uc.verifySorted(sortedList),
			Analysis:    pkg.CalculateAnalysis(count, duration),
			IsArray:     false,
		}
	}
}

// CustomRandomSort generates and sorts a random list of specified size
func (uc *UseCase) CustomRandomSort(algorithmName string, count int) SortResult {
	startTime := time.Now()

	if uc.algorithmUsesArray(algorithmName) {
		// Array-based algorithms
		numbers := uc.generator.GenerateIntSliceDefault(count)
		sortedArray := uc.executeSortArray(algorithmName, numbers)
		duration := time.Since(startTime)

		return SortResult{
			SortedList:  nil,
			SortedArray: sortedArray,
			Duration:    duration,
			Count:       count,
			IsSorted:    pkg.IsSortedSlice(sortedArray),
			Analysis:    pkg.CalculateAnalysis(count, duration),
			IsArray:     true,
		}
	} else {
		// Linked list algorithms
		head := uc.generateRandomList(count)
		sortedList := uc.executeSortList(algorithmName, head)
		duration := time.Since(startTime)

		return SortResult{
			SortedList:  sortedList,
			SortedArray: nil,
			Duration:    duration,
			Count:       count,
			IsSorted:    uc.verifySorted(sortedList),
			Analysis:    pkg.CalculateAnalysis(count, duration),
			IsArray:     false,
		}
	}
}

// BenchmarkSort runs a benchmark with predefined size
func (uc *UseCase) BenchmarkSort(algorithmName string, count int) SortResult {
	return uc.CustomRandomSort(algorithmName, count)
}

// RunAllBenchmarks executes all predefined benchmarks for an algorithm
func (uc *UseCase) RunAllBenchmarks(algorithmName string) BenchmarkSummary {
	sizes := []int{500, 1000, 5000, 10000}
	results := make([]BenchmarkResult, 0, len(sizes))

	for _, count := range sizes {
		result := uc.BenchmarkSort(algorithmName, count)
		results = append(results, BenchmarkResult{
			Count:    count,
			Duration: result.Duration,
			IsSorted: result.IsSorted,
		})
	}

	scaling := pkg.CalculateScaling(results)

	return BenchmarkSummary{
		Results: results,
		Scaling: scaling,
	}
}

// GetListPartial returns the first n elements of a list as a slice
func (uc *UseCase) GetListPartial(head *merge_sort.Node, maxCount int) []int {
	var result []int
	current := head
	count := 0

	for current != nil && count < maxCount {
		result = append(result, current.Value)
		current = current.Next
		count++
	}

	return result
}

// algorithmUsesArray determines if an algorithm works with arrays or linked lists
func (uc *UseCase) algorithmUsesArray(algorithmName string) bool {
	switch algorithmName {
	case "Quick Sort", "Heap Sort", "Bubble Sort", "Insertion Sort":
		return true
	case "Merge Sort":
		return false
	default:
		return false // Default to linked list
	}
}

// executeSortArray dispatches to array-based sorting algorithms
func (uc *UseCase) executeSortArray(algorithmName string, arr []int) []int {
	switch algorithmName {
	case "Quick Sort":
		return quick_sort.QuickSort(arr)
	case "Bubble Sort":
		return bubble_sort.BubbleSortOptimized(arr)
	default:
		// For future array algorithms, we'll add cases here
		return quick_sort.QuickSort(arr) // Default fallback
	}
}

// executeSortList dispatches to linked list sorting algorithms
func (uc *UseCase) executeSortList(algorithmName string, head *merge_sort.Node) *merge_sort.Node {
	switch algorithmName {
	case "Merge Sort":
		return merge_sort.MergeSort(head)
	default:
		// For future linked list algorithms, we'll add cases here
		return merge_sort.MergeSort(head) // Default fallback
	}
}

// createListFromSlice creates a linked list from a slice of integers
func (uc *UseCase) createListFromSlice(values []int) *merge_sort.Node {
	if len(values) == 0 {
		return nil
	}

	head := &merge_sort.Node{Value: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		newNode := &merge_sort.Node{Value: values[i]}
		current.Next = newNode
		current = newNode
	}

	return head
}

// generateRandomList creates a linked list with random numbers from 1 to 1000
func (uc *UseCase) generateRandomList(count int) *merge_sort.Node {
	if count <= 0 {
		return nil
	}

	numbers := uc.generator.GenerateIntSliceDefault(count)
	return uc.createListFromSlice(numbers)
}

// verifySorted checks if a linked list is correctly sorted
func (uc *UseCase) verifySorted(head *merge_sort.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	current := head
	for current.Next != nil {
		if current.Value > current.Next.Value {
			return false
		}
		current = current.Next
	}

	return true
}
