package mergesort

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// RunBenchmarkTerminal provides an advanced interface for testing MergeSort performance
func RunBenchmarkTerminal() {
	fmt.Println("\n\n[   Merge Sort - Advanced Testing   ]")
	fmt.Println("Choose a testing option:")
	fmt.Println()
	fmt.Println("1. Manual input (enter numbers manually)")
	fmt.Println("2. Custom random list (specify quantity)")
	fmt.Println("3. Benchmark 500 random numbers")
	fmt.Println("4. Benchmark 1,000 random numbers")
	fmt.Println("5. Benchmark 5,000 random numbers")
	fmt.Println("6. Benchmark 10,000 random numbers")
	fmt.Println("7. Run all benchmarks")
	fmt.Println("8. Back to main menu")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice (1-8): ")
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		runManualInput()
	case "2":
		runCustomRandom()
	case "3":
		runBenchmark(500)
	case "4":
		runBenchmark(1000)
	case "5":
		runBenchmark(5000)
	case "6":
		runBenchmark(10000)
	case "7":
		runAllBenchmarks()
	case "8":
		return
	default:
		fmt.Println("Invalid choice. Please select a valid option (1-8).")
		return // Exit instead of recursive call
	}
}

// runManualInput handles manual number input (original functionality)
func runManualInput() {
	fmt.Println("\n=== Manual Input Mode ===")
	fmt.Println("Enter numbers one by one and press Enter. To stop and sort, just press Enter on an empty line.")

	var head *Node
	var tail *Node

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		newNode := &Node{Value: num}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	if head == nil {
		fmt.Println("The list is empty. Nothing to sort.")
		return
	}

	fmt.Println("\nOriginal List:")
	printList(head)

	// Measure sorting time
	startTime := time.Now()
	sortedList := MergeSort(head)
	duration := time.Since(startTime)

	fmt.Println("Sorted List:")
	printList(sortedList)

	fmt.Printf("\n✅ Sorting completed in: %v\n", duration)
	fmt.Printf("📊 Performance: %d numbers sorted\n", getListLength(sortedList))
}

// runCustomRandom generates a custom-sized random list
func runCustomRandom() {
	fmt.Println("\n=== Custom Random List Mode ===")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of random numbers to generate (1-1,000,000): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	count, err := strconv.Atoi(input)
	if err != nil || count <= 0 || count > 1000000 {
		fmt.Println("Invalid input. Please enter a number between 1 and 1,000,000.")
		return
	}

	fmt.Printf("\n🎲 Generating %d random numbers...\n", count)
	head := generateRandomList(count)

	fmt.Printf("📝 Random list with %d numbers generated successfully!\n", count)
	fmt.Println("🔄 Starting sort...")

	// Measure sorting time
	startTime := time.Now()
	sortedList := MergeSort(head)
	duration := time.Since(startTime)

	fmt.Printf("\n✅ Sorting completed in: %v\n", duration)
	fmt.Printf("📊 Performance: %d numbers sorted\n", count)
	fmt.Printf("⚡ Average time per number: %.2f μs\n", float64(duration.Nanoseconds())/float64(count)/1000.0)

	// Ask if user wants to see the sorted list (for smaller lists)
	if count <= 50 {
		fmt.Print("\nDo you want to see the sorted list? (y/n): ")
		response, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(response)) == "y" {
			fmt.Println("\nSorted List:")
			printList(sortedList)
		}
	} else if count <= 1000 {
		fmt.Print("\nDo you want to see the first 20 numbers of the sorted list? (y/n): ")
		response, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(response)) == "y" {
			fmt.Println("\nFirst 20 numbers of sorted list:")
			printListPartial(sortedList, 20)
		}
	}
}

// runBenchmark runs a benchmark with a predefined number of elements
func runBenchmark(count int) {
	fmt.Printf("\n=== Benchmark: %s Random Numbers ===\n", formatNumber(count))
	
	fmt.Printf("🎲 Generating %s random numbers...\n", formatNumber(count))
	head := generateRandomList(count)

	fmt.Printf("📝 List with %s numbers generated successfully!\n", formatNumber(count))
	fmt.Println("🔄 Starting sort...")

	// Measure sorting time
	startTime := time.Now()
	sortedList := MergeSort(head)
	duration := time.Since(startTime)

	// Verify the list is sorted (optional verification)
	isSorted := verifySorted(sortedList)
	
	fmt.Printf("\n✅ Sorting completed in: %v\n", duration)
	fmt.Printf("📊 Performance: %s numbers sorted\n", formatNumber(count))
	fmt.Printf("⚡ Average time per number: %.2f μs\n", float64(duration.Nanoseconds())/float64(count)/1000.0)
	
	if isSorted {
		fmt.Println("✅ Verification: List is correctly sorted!")
	} else {
		fmt.Println("❌ Warning: List may not be correctly sorted!")
	}

	// Performance analysis
	analyzePerformance(count, duration)
}

// runAllBenchmarks runs all predefined benchmarks
func runAllBenchmarks() {
	fmt.Println("\n=== Running All Benchmarks ===")
	fmt.Println("This will test MergeSort performance with different input sizes...")
	fmt.Println()

	benchmarks := []int{500, 1000, 5000, 10000}
	results := make([]BenchmarkResult, 0, len(benchmarks))

	for i, count := range benchmarks {
		fmt.Printf("[%d/%d] ", i+1, len(benchmarks))
		
		// Generate list
		fmt.Printf("Generating %s numbers... ", formatNumber(count))
		head := generateRandomList(count)
		
		// Sort and measure
		fmt.Print("Sorting... ")
		startTime := time.Now()
		sortedList := MergeSort(head)
		duration := time.Since(startTime)
		
		// Verify
		isSorted := verifySorted(sortedList)
		
		result := BenchmarkResult{
			Count:    count,
			Duration: duration,
			IsSorted: isSorted,
		}
		results = append(results, result)
		
		fmt.Printf("Done in %v ✅\n", duration)
	}

	// Display results summary
	displayBenchmarkSummary(results)
}

// BenchmarkResult stores the results of a benchmark run
type BenchmarkResult struct {
	Count    int
	Duration time.Duration
	IsSorted bool
}

// generateRandomList creates a linked list with random numbers from 1 to 1000
func generateRandomList(count int) *Node {
	if count <= 0 {
		return nil
	}

	// Seed the random number generator
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

// getListLength returns the length of a linked list
func getListLength(head *Node) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// printListPartial prints only the first n elements of a linked list
func printListPartial(node *Node, maxCount int) {
	current := node
	count := 0
	for current != nil && count < maxCount {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
		count++
	}
	if current != nil {
		fmt.Println("... (and more)")
	} else {
		fmt.Println("nil")
	}
}

// verifySorted checks if a linked list is correctly sorted
func verifySorted(head *Node) bool {
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

// formatNumber formats large numbers with commas for readability
func formatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}
	
	str := fmt.Sprintf("%d", n)
	result := ""
	
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}
	
	return result
}

// analyzePerformance provides performance analysis based on theoretical complexity
func analyzePerformance(count int, actualDuration time.Duration) {
	// Theoretical O(n log n) comparison
	theoreticalOps := float64(count) * math.Log2(float64(count))
	actualNanoseconds := float64(actualDuration.Nanoseconds())
	
	fmt.Printf("📈 Theoretical O(n log n): %.0f operations\n", theoreticalOps)
	fmt.Printf("⏱️  Time per operation: %.2f ns\n", actualNanoseconds/theoreticalOps)
}

// displayBenchmarkSummary shows a comprehensive summary of all benchmark results
func displayBenchmarkSummary(results []BenchmarkResult) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("               BENCHMARK SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	
	fmt.Printf("%-12s %-15s %-15s %-10s\n", "Count", "Time", "μs/number", "Status")
	fmt.Println(strings.Repeat("-", 60))
	
	for _, result := range results {
		timePerNumber := float64(result.Duration.Nanoseconds()) / float64(result.Count) / 1000.0
		status := "✅"
		if !result.IsSorted {
			status = "❌"
		}
		
		fmt.Printf("%-12s %-15v %-15.2f %-10s\n", 
			formatNumber(result.Count), 
			result.Duration, 
			timePerNumber, 
			status)
	}
	
	fmt.Println(strings.Repeat("=", 60))
	
	// Calculate and display scaling analysis
	if len(results) > 1 {
		fmt.Println("\n📊 SCALING ANALYSIS:")
		for i := 1; i < len(results); i++ {
			prev := results[i-1]
			curr := results[i]
			
			sizeRatio := float64(curr.Count) / float64(prev.Count)
			timeRatio := float64(curr.Duration.Nanoseconds()) / float64(prev.Duration.Nanoseconds())
			
			fmt.Printf("   %s → %s: %.2fx size, %.2fx time\n", 
				formatNumber(prev.Count), 
				formatNumber(curr.Count), 
				sizeRatio, 
				timeRatio)
		}
		
		fmt.Println("\n💡 Note: MergeSort has O(n log n) complexity.")
		fmt.Println("   Expected time ratio for 2x size: ~2.2x time")
		fmt.Println("   Expected time ratio for 10x size: ~33x time")
	}
}
