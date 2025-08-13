package sorting

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"
)

// Terminal handles all user interface interactions for sorting algorithms
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

// RunSortingInterface provides the main interface for sorting algorithm testing
func RunSortingInterface() {
	terminal := NewTerminal()
	terminal.showSortingMenu()
}

func (t *Terminal) showSortingMenu() {
	fmt.Println("\n\n[   Sorting Algorithms - Advanced Testing   ]")
	fmt.Println("Choose a sorting algorithm:")
	fmt.Println()
	fmt.Println("1. Merge Sort")
	fmt.Println("2. Quick Sort (Coming Soon)")
	fmt.Println("3. Heap Sort (Coming Soon)")
	fmt.Println("4. Bubble Sort (Coming Soon)")
	fmt.Println("5. Insertion Sort (Coming Soon)")
	fmt.Println("6. Back to main menu")
	fmt.Println()

	choice := t.getMenuChoice("Enter your choice (1-6): ")

	switch choice {
	case "1":
		t.showAlgorithmMenu("Merge Sort")
	case "2":
		t.showAlgorithmMenu("Quick Sort")
	case "3", "4", "5":
		fmt.Println("This algorithm is coming soon! Stay tuned. 🚀")
		t.showSortingMenu()
	case "6":
		return
	default:
		fmt.Println("Invalid choice. Please select a valid option (1-6).")
		t.showSortingMenu()
	}
}

func (t *Terminal) showAlgorithmMenu(algorithmName string) {
	fmt.Printf("\n\n[   %s - Advanced Testing   ]\n", algorithmName)
	fmt.Println("Choose a testing option:")
	fmt.Println()
	fmt.Println("1. Manual input (enter numbers manually)")
	fmt.Println("2. Custom random list (specify quantity)")
	fmt.Println("3. Benchmark 500 random numbers")
	fmt.Println("4. Benchmark 1,000 random numbers")
	fmt.Println("5. Benchmark 5,000 random numbers")
	fmt.Println("6. Benchmark 10,000 random numbers")
	fmt.Println("7. Run all benchmarks")
	fmt.Println("8. Back to sorting menu")
	fmt.Println()

	choice := t.getMenuChoice("Enter your choice (1-8): ")

	switch choice {
	case "1":
		t.runManualInput(algorithmName)
	case "2":
		t.runCustomRandom(algorithmName)
	case "3":
		t.runBenchmark(algorithmName, 500)
	case "4":
		t.runBenchmark(algorithmName, 1000)
	case "5":
		t.runBenchmark(algorithmName, 5000)
	case "6":
		t.runBenchmark(algorithmName, 10000)
	case "7":
		t.runAllBenchmarks(algorithmName)
	case "8":
		t.showSortingMenu()
	default:
		fmt.Println("Invalid choice. Please select a valid option (1-8).")
		t.showAlgorithmMenu(algorithmName)
	}
}

func (t *Terminal) runManualInput(algorithmName string) {
	fmt.Printf("\n=== %s - Manual Input Mode ===\n", algorithmName)
	fmt.Println("Enter numbers one by one and press Enter. To stop and sort, just press Enter on an empty line.")

	numbers := t.collectNumbers()
	if len(numbers) == 0 {
		fmt.Println("The list is empty. Nothing to sort.")
		return
	}

	result := t.useCase.ManualSort(algorithmName, numbers)

	if result.IsArray {
		fmt.Println("\nOriginal Array:")
		t.printSlice(numbers)

		fmt.Println("Sorted Array:")
		t.printSlice(result.SortedArray)
	} else {
		fmt.Println("\nOriginal List:")
		t.printSlice(numbers)

		fmt.Println("Sorted List:")
		t.printList(result.SortedList)
	}

	t.printPerformanceInfo(result)
}

func (t *Terminal) runCustomRandom(algorithmName string) {
	fmt.Printf("\n=== %s - Custom Random List Mode ===\n", algorithmName)
	
	count := t.getNumber("Enter the number of random numbers to generate (1-1,000,000): ", 1, 1000000)
	if count == -1 {
		return
	}

	fmt.Printf("\n🎲 Generating %s random numbers...\n", t.formatNumber(count))
	
	result := t.useCase.CustomRandomSort(algorithmName, count)

	fmt.Printf("📝 Random list with %s numbers generated successfully!\n", t.formatNumber(count))
	fmt.Println("🔄 Starting sort...")

	t.printPerformanceInfo(result)
	t.askToShowList(result, count)
}

func (t *Terminal) runBenchmark(algorithmName string, count int) {
	fmt.Printf("\n=== %s Benchmark: %s Random Numbers ===\n", algorithmName, t.formatNumber(count))
	
	fmt.Printf("🎲 Generating %s random numbers...\n", t.formatNumber(count))
	
	result := t.useCase.BenchmarkSort(algorithmName, count)

	fmt.Printf("📝 List with %s numbers generated successfully!\n", t.formatNumber(count))
	fmt.Println("🔄 Starting sort...")

	t.printDetailedPerformance(result)
}

func (t *Terminal) runAllBenchmarks(algorithmName string) {
	fmt.Printf("\n=== %s - Running All Benchmarks ===\n", algorithmName)
	fmt.Println("This will test sorting performance with different input sizes...")
	fmt.Println()

	summary := t.useCase.RunAllBenchmarks(algorithmName)
	
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
			if result.IsArray {
				t.printSlice(result.SortedArray)
			} else {
				t.printList(result.SortedList)
			}
		}
	} else if count <= 1000 {
		if t.askYesNo("\nDo you want to see the first 20 numbers of the sorted list? (y/n): ") {
			fmt.Println("\nFirst 20 numbers of sorted list:")
			if result.IsArray {
				t.printSlicePartial(result.SortedArray, 20)
			} else {
				t.printListPartial(result.SortedList, 20)
			}
		}
	}
}

func (t *Terminal) askYesNo(prompt string) bool {
	fmt.Print(prompt)
	response, _ := t.reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(response)) == "y"
}

func (t *Terminal) printSlice(numbers []int) {
	fmt.Print("[")
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println("]")
}

func (t *Terminal) printList(node *merge_sort.Node) {
	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (t *Terminal) printListPartial(node *merge_sort.Node, maxCount int) {
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

func (t *Terminal) printSlicePartial(arr []int, maxCount int) {
	count := maxCount
	if len(arr) < maxCount {
		count = len(arr)
	}
	
	fmt.Print("[")
	for i := 0; i < count; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(arr[i])
	}
	
	if len(arr) > maxCount {
		fmt.Println(", ... (and more)]")
	} else {
		fmt.Println("]")
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
		
		fmt.Println("\n💡 Note: Performance depends on the algorithm complexity.")
		fmt.Println("   For O(n log n) algorithms: Expected time ratio for 2x size: ~2.2x time")
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
