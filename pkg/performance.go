package pkg

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// PerformanceAnalysis contains performance metrics
type PerformanceAnalysis struct {
	TimePerNumber       float64 // microseconds per number
	TheoreticalOps      float64 // O(n log n) operations
	TimePerOperation    float64 // nanoseconds per operation
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

// BenchmarkSummary contains results from multiple benchmarks
type BenchmarkSummary struct {
	Results []BenchmarkResult
	Scaling []ScalingAnalysis
}

// CalculateAnalysis provides performance analysis based on theoretical complexity
func CalculateAnalysis(count int, duration time.Duration) PerformanceAnalysis {
	timePerNumber := float64(duration.Nanoseconds()) / float64(count) / 1000.0 // microseconds
	theoreticalOps := float64(count) * math.Log2(float64(count))
	timePerOperation := float64(duration.Nanoseconds()) / theoreticalOps

	return PerformanceAnalysis{
		TimePerNumber:    timePerNumber,
		TheoreticalOps:   theoreticalOps,
		TimePerOperation: timePerOperation,
	}
}

// CalculateScaling analyzes performance scaling between different input sizes
func CalculateScaling(results []BenchmarkResult) []ScalingAnalysis {
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

// PrintPerformanceInfo prints basic performance information
func PrintPerformanceInfo(count int, duration time.Duration, analysis PerformanceAnalysis) {
	fmt.Printf("\n✅ Sorting completed in: %v\n", duration)
	fmt.Printf("📊 Performance: %d numbers sorted\n", count)
	fmt.Printf("⚡ Average time per number: %.2f μs\n", analysis.TimePerNumber)
}

// PrintDetailedPerformance prints detailed performance analysis
func PrintDetailedPerformance(count int, duration time.Duration, analysis PerformanceAnalysis, isSorted bool) {
	PrintPerformanceInfo(count, duration, analysis)
	
	if isSorted {
		fmt.Println("✅ Verification: List is correctly sorted!")
	} else {
		fmt.Println("❌ Warning: List may not be correctly sorted!")
	}

	fmt.Printf("📈 Theoretical O(n log n): %.0f operations\n", analysis.TheoreticalOps)
	fmt.Printf("⏱️  Time per operation: %.2f ns\n", analysis.TimePerOperation)
}

// PrintBenchmarkSummary displays a comprehensive benchmark summary
func PrintBenchmarkSummary(summary BenchmarkSummary) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("               BENCHMARK SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	
	fmt.Printf("%-12s %-15s %-15s %-10s\n", "Count", "Time", "μs/number", "Status")
	fmt.Println(strings.Repeat("-", 60))
	
	for _, result := range summary.Results {
		timePerNumber := float64(result.Duration.Nanoseconds()) / float64(result.Count) / 1000.0
		status := "✅"
		if !result.IsSorted {
			status = "❌"
		}
		
		fmt.Printf("%-12s %-15v %-15.2f %-10s\n", 
			FormatNumber(result.Count), 
			result.Duration, 
			timePerNumber, 
			status)
	}
	
	fmt.Println(strings.Repeat("=", 60))
	
	if len(summary.Scaling) > 0 {
		fmt.Println("\n📊 SCALING ANALYSIS:")
		for _, scale := range summary.Scaling {
			fmt.Printf("   %s → %s: %.2fx size, %.2fx time\n", 
				FormatNumber(scale.FromSize), 
				FormatNumber(scale.ToSize), 
				scale.SizeRatio, 
				scale.TimeRatio)
		}
		
		fmt.Println("\n💡 Note: Performance depends on the algorithm complexity.")
		fmt.Println("   For O(n log n) algorithms: Expected time ratio for 2x size: ~2.2x time")
	}
}
