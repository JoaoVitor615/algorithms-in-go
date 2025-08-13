package mergesort

import (
	"math"
	"math/rand"
	"time"
)

// UseCase represents the business logic layer for MergeSort operations
type UseCase struct{}

// NewUseCase creates a new UseCase instance
func NewUseCase() *UseCase {
	return &UseCase{}
}

// SortResult contains the result of a sorting operation
type SortResult struct {
	SortedList   *Node
	Duration     time.Duration
	Count        int
	IsSorted     bool
	Analysis     PerformanceAnalysis
}

// PerformanceAnalysis contains performance metrics
type PerformanceAnalysis struct {
	TimePerNumber       float64 // microseconds per number
	TheoreticalOps      float64 // O(n log n) operations
	TimePerOperation    float64 // nanoseconds per operation
}

// BenchmarkSummary contains results from multiple benchmarks
type BenchmarkSummary struct {
	Results []BenchmarkResult
	Scaling []ScalingAnalysis
}

// BenchmarkResult stores individual benchmark results
type BenchmarkResult struct {
	Count    int
	Duration time.Duration
	IsSorted bool
}

// ScalingAnalysis compares performance between different input sizes
type ScalingAnalysis struct {
	FromSize   int
	ToSize     int
	SizeRatio  float64
	TimeRatio  float64
}

// ManualSort sorts a manually created linked list
func (uc *UseCase) ManualSort(head *Node) SortResult {
	if head == nil {
		return SortResult{
			SortedList: nil,
			Duration:   0,
			Count:      0,
			IsSorted:   true,
			Analysis:   PerformanceAnalysis{},
		}
	}

	count := uc.getListLength(head)
	startTime := time.Now()
	sortedList := MergeSort(head)
	duration := time.Since(startTime)

	return SortResult{
		SortedList: sortedList,
		Duration:   duration,
		Count:      count,
		IsSorted:   uc.verifySorted(sortedList),
		Analysis:   uc.calculateAnalysis(count, duration),
	}
}

// CustomRandomSort generates and sorts a random list of specified size
func (uc *UseCase) CustomRandomSort(count int) SortResult {
	head := uc.generateRandomList(count)
	startTime := time.Now()
	sortedList := MergeSort(head)
	duration := time.Since(startTime)

	return SortResult{
		SortedList: sortedList,
		Duration:   duration,
		Count:      count,
		IsSorted:   uc.verifySorted(sortedList),
		Analysis:   uc.calculateAnalysis(count, duration),
	}
}

// BenchmarkSort runs a benchmark with predefined size
func (uc *UseCase) BenchmarkSort(count int) SortResult {
	return uc.CustomRandomSort(count)
}

// RunAllBenchmarks executes all predefined benchmarks
func (uc *UseCase) RunAllBenchmarks() BenchmarkSummary {
	sizes := []int{500, 1000, 5000, 10000}
	results := make([]BenchmarkResult, 0, len(sizes))

	for _, count := range sizes {
		result := uc.BenchmarkSort(count)
		results = append(results, BenchmarkResult{
			Count:    count,
			Duration: result.Duration,
			IsSorted: result.IsSorted,
		})
	}

	scaling := uc.calculateScaling(results)

	return BenchmarkSummary{
		Results: results,
		Scaling: scaling,
	}
}

// CreateListFromSlice creates a linked list from a slice of integers
func (uc *UseCase) CreateListFromSlice(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{Value: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		newNode := &Node{Value: values[i]}
		current.Next = newNode
		current = newNode
	}

	return head
}

// ListToSlice converts a linked list to a slice
func (uc *UseCase) ListToSlice(head *Node) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// GetListPartial returns the first n elements of a list as a slice
func (uc *UseCase) GetListPartial(head *Node, maxCount int) []int {
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

// Private helper methods

func (uc *UseCase) generateRandomList(count int) *Node {
	if count <= 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	head := &Node{Value: rand.Intn(1000) + 1}
	current := head

	for i := 1; i < count; i++ {
		newNode := &Node{Value: rand.Intn(1000) + 1}
		current.Next = newNode
		current = newNode
	}

	return head
}

func (uc *UseCase) getListLength(head *Node) int {
	count := 0
	current := head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}

func (uc *UseCase) verifySorted(head *Node) bool {
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

func (uc *UseCase) calculateAnalysis(count int, duration time.Duration) PerformanceAnalysis {
	timePerNumber := float64(duration.Nanoseconds()) / float64(count) / 1000.0 // microseconds
	theoreticalOps := float64(count) * math.Log2(float64(count))
	timePerOperation := float64(duration.Nanoseconds()) / theoreticalOps

	return PerformanceAnalysis{
		TimePerNumber:    timePerNumber,
		TheoreticalOps:   theoreticalOps,
		TimePerOperation: timePerOperation,
	}
}

func (uc *UseCase) calculateScaling(results []BenchmarkResult) []ScalingAnalysis {
	if len(results) < 2 {
		return nil
	}

	scaling := make([]ScalingAnalysis, 0, len(results)-1)

	for i := 1; i < len(results); i++ {
		prev := results[i-1]
		curr := results[i]

		sizeRatio := float64(curr.Count) / float64(prev.Count)
		timeRatio := float64(curr.Duration.Nanoseconds()) / float64(prev.Duration.Nanoseconds())

		scaling = append(scaling, ScalingAnalysis{
			FromSize:  prev.Count,
			ToSize:    curr.Count,
			SizeRatio: sizeRatio,
			TimeRatio: timeRatio,
		})
	}

	return scaling
}
