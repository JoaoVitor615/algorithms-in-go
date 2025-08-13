package mergesort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Terminal handles all user interface interactions
type Terminal struct {
	useCase *UseCase
	reader  *bufio.Reader
}

// NewTerminal creates a new Terminal instance
func NewTerminal() *Terminal {
	return &Terminal{
		useCase: NewUseCase(),
		reader:  bufio.NewReader(os.Stdin),
	}
}

// RunAdvancedInterface provides the main interface for MergeSort testing
func RunAdvancedInterface() {
	terminal := NewTerminal()
	terminal.showMainMenu()
}

// RunTerminal provides the original simple interface (legacy compatibility)
func RunTerminal() {
	terminal := NewTerminal()
	terminal.runManualInput()
}

func (t *Terminal) showMainMenu() {
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

	choice := t.getMenuChoice("Enter your choice (1-8): ")

	switch choice {
	case "1":
		t.runManualInput()
	case "2":
		t.runCustomRandom()
	case "3":
		t.runBenchmark(500)
	case "4":
		t.runBenchmark(1000)
	case "5":
		t.runBenchmark(5000)
	case "6":
		t.runBenchmark(10000)
	case "7":
		t.runAllBenchmarks()
	case "8":
		return
	default:
		fmt.Println("Invalid choice. Please select a valid option (1-8).")
	}
}

func (t *Terminal) runManualInput() {
	fmt.Println("\n=== Manual Input Mode ===")
	fmt.Println("Enter numbers one by one and press Enter. To stop and sort, just press Enter on an empty line.")

	numbers := t.collectNumbers()
	if len(numbers) == 0 {
		fmt.Println("The list is empty. Nothing to sort.")
		return
	}

	head := t.useCase.CreateListFromSlice(numbers)
	
	fmt.Println("\nOriginal List:")
	t.printList(head)

	result := t.useCase.ManualSort(head)

	fmt.Println("Sorted List:")
	t.printList(result.SortedList)

	t.printPerformanceInfo(result)
}

func (t *Terminal) runCustomRandom() {
	fmt.Println("\n=== Custom Random List Mode ===")
	
	count := t.getNumber("Enter the number of random numbers to generate (1-1,000,000): ", 1, 1000000)
	if count == -1 {
		return
	}

	fmt.Printf("\n🎲 Generating %s random numbers...\n", t.formatNumber(count))
	
	result := t.useCase.CustomRandomSort(count)

	fmt.Printf("📝 Random list with %s numbers generated successfully!\n", t.formatNumber(count))
	fmt.Println("🔄 Starting sort...")

	t.printPerformanceInfo(result)
	t.askToShowList(result, count)
}

func (t *Terminal) runBenchmark(count int) {
	fmt.Printf("\n=== Benchmark: %s Random Numbers ===\n", t.formatNumber(count))
	
	fmt.Printf("🎲 Generating %s random numbers...\n", t.formatNumber(count))
	
	result := t.useCase.BenchmarkSort(count)

	fmt.Printf("📝 List with %s numbers generated successfully!\n", t.formatNumber(count))
	fmt.Println("🔄 Starting sort...")

	t.printDetailedPerformance(result)
}

func (t *Terminal) runAllBenchmarks() {
	fmt.Println("\n=== Running All Benchmarks ===")
	fmt.Println("This will test MergeSort performance with different input sizes...")
	fmt.Println()

	summary := t.useCase.RunAllBenchmarks()
	
	for i, result := range summary.Results {
		fmt.Printf("[%d/%d] Generating %s numbers... Sorting... Done in %v ✅\n", 
			i+1, len(summary.Results), t.formatNumber(result.Count), result.Duration)
	}

	t.displayBenchmarkSummary(summary)
}

func (t *Terminal) collectNumbers() []int {
	var numbers []int

	for {
		fmt.Print("Enter a number: ")
		input, _ := t.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func (t *Terminal) getMenuChoice(prompt string) string {
	fmt.Print(prompt)
	input, _ := t.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (t *Terminal) getNumber(prompt string, min, max int) int {
	fmt.Print(prompt)
	input, _ := t.reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil || num < min || num > max {
		fmt.Printf("Invalid input. Please enter a number between %s and %s.\n", 
			t.formatNumber(min), t.formatNumber(max))
		return -1
	}

	return num
}

func (t *Terminal) askToShowList(result SortResult, count int) {
	if count <= 50 {
		if t.askYesNo("\nDo you want to see the sorted list? (y/n): ") {
			fmt.Println("\nSorted List:")
			t.printList(result.SortedList)
		}
	} else if count <= 1000 {
		if t.askYesNo("\nDo you want to see the first 20 numbers of the sorted list? (y/n): ") {
			fmt.Println("\nFirst 20 numbers of sorted list:")
			t.printListPartial(result.SortedList, 20)
		}
	}
}

func (t *Terminal) askYesNo(prompt string) bool {
	fmt.Print(prompt)
	response, _ := t.reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(response)) == "y"
}

func (t *Terminal) printList(node *Node) {
	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (t *Terminal) printListPartial(node *Node, maxCount int) {
	partial := t.useCase.GetListPartial(node, maxCount)
	
	for _, value := range partial {
		fmt.Printf("%d -> ", value)
	}
	
	if len(partial) == maxCount {
		fmt.Println("... (and more)")
	} else {
		fmt.Println("nil")
	}
}

func (t *Terminal) printPerformanceInfo(result SortResult) {
	fmt.Printf("\n✅ Sorting completed in: %v\n", result.Duration)
	fmt.Printf("📊 Performance: %d numbers sorted\n", result.Count)
	fmt.Printf("⚡ Average time per number: %.2f μs\n", result.Analysis.TimePerNumber)
}

func (t *Terminal) printDetailedPerformance(result SortResult) {
	t.printPerformanceInfo(result)
	
	if result.IsSorted {
		fmt.Println("✅ Verification: List is correctly sorted!")
	} else {
		fmt.Println("❌ Warning: List may not be correctly sorted!")
	}

	fmt.Printf("📈 Theoretical O(n log n): %.0f operations\n", result.Analysis.TheoreticalOps)
	fmt.Printf("⏱️  Time per operation: %.2f ns\n", result.Analysis.TimePerOperation)
}

func (t *Terminal) displayBenchmarkSummary(summary BenchmarkSummary) {
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
			t.formatNumber(result.Count), 
			result.Duration, 
			timePerNumber, 
			status)
	}
	
	fmt.Println(strings.Repeat("=", 60))
	
	if len(summary.Scaling) > 0 {
		fmt.Println("\n📊 SCALING ANALYSIS:")
		for _, scale := range summary.Scaling {
			fmt.Printf("   %s → %s: %.2fx size, %.2fx time\n", 
				t.formatNumber(scale.FromSize), 
				t.formatNumber(scale.ToSize), 
				scale.SizeRatio, 
				scale.TimeRatio)
		}
		
		fmt.Println("\n💡 Note: MergeSort has O(n log n) complexity.")
		fmt.Println("   Expected time ratio for 2x size: ~2.2x time")
		fmt.Println("   Expected time ratio for 10x size: ~33x time")
	}
}

func (t *Terminal) formatNumber(n int) string {
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