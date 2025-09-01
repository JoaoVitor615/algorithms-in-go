package sorting

import (
	"fmt"

	"github.com/JoaoVitor615/algorithms-in-go/pkg"
	"github.com/JoaoVitor615/algorithms-in-go/sorting/merge_sort"
)

// Terminal handles all user interface interactions for sorting algorithms
type Terminal struct {
	useCase *UseCase
	input   *pkg.InputReader
}

// NewTerminal creates a new Terminal instance
func NewTerminal() *Terminal {
	return &Terminal{
		useCase: NewUseCase(),
		input:   pkg.NewInputReader(),
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
	fmt.Println("1. Merge Sort")
	fmt.Println("2. Quick Sort")
	fmt.Println("3. Bubble Sort")
	fmt.Println("4. Heap Sort (Coming Soon)")
	fmt.Println("5. Insertion Sort (Coming Soon)")
	fmt.Println("6. Back to main menu")
	fmt.Println()

	choice := t.getMenuChoice("Enter your choice (1-6): ")

	switch choice {
	case "1":
		t.showAlgorithmMenu("Merge Sort")
	case "2":
		t.showAlgorithmMenu("Quick Sort")
	case "3":
		t.showAlgorithmMenu("Bubble Sort")
	case "4", "5":
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
	pkg.PrintSubHeader(fmt.Sprintf("%s - Manual Input Mode", algorithmName))
	fmt.Println("Enter numbers one by one and press Enter. To stop and sort, just press Enter on an empty line.")

	numbers := t.input.ReadNumbers()
	if pkg.IsEmpty(numbers) {
		fmt.Println("The list is empty. Nothing to sort.")
		return
	}

	result := t.useCase.ManualSort(algorithmName, numbers)

	if result.IsArray {
		fmt.Println("\nOriginal Array:")
		pkg.PrintSlice(numbers)

		fmt.Println("Sorted Array:")
		pkg.PrintSlice(result.SortedArray)
	} else {
		fmt.Println("\nOriginal List:")
		pkg.PrintSlice(numbers)

		fmt.Println("Sorted List:")
		t.printList(result.SortedList)
	}

	pkg.PrintPerformanceInfo(result.Count, result.Duration, result.Analysis)
}

func (t *Terminal) runCustomRandom(algorithmName string) {
	pkg.PrintSubHeader(fmt.Sprintf("%s - Custom Random List Mode", algorithmName))
	
	count := t.input.ReadIntOrDefault("Enter the number of random numbers to generate (1-1,000,000): ", 1, 1000000)
	if count == -1 {
		fmt.Printf("Invalid input. Please enter a number between %s and %s.\n", 
			pkg.FormatNumber(1), pkg.FormatNumber(1000000))
		return
	}

	fmt.Printf("\n🎲 Generating %s random numbers...\n", pkg.FormatNumber(count))
	
	result := t.useCase.CustomRandomSort(algorithmName, count)

	fmt.Printf("📝 Random list with %s numbers generated successfully!\n", pkg.FormatNumber(count))
	fmt.Println("🔄 Starting sort...")

	pkg.PrintPerformanceInfo(result.Count, result.Duration, result.Analysis)
	t.askToShowList(result, count)
}

func (t *Terminal) runBenchmark(algorithmName string, count int) {
	pkg.PrintSubHeader(fmt.Sprintf("%s Benchmark: %s Random Numbers", algorithmName, pkg.FormatNumber(count)))
	
	fmt.Printf("🎲 Generating %s random numbers...\n", pkg.FormatNumber(count))
	
	result := t.useCase.BenchmarkSort(algorithmName, count)

	fmt.Printf("📝 List with %s numbers generated successfully!\n", pkg.FormatNumber(count))
	fmt.Println("🔄 Starting sort...")

	pkg.PrintDetailedPerformance(result.Count, result.Duration, result.Analysis, result.IsSorted)
}

func (t *Terminal) runAllBenchmarks(algorithmName string) {
	pkg.PrintSubHeader(fmt.Sprintf("%s - Running All Benchmarks", algorithmName))
	fmt.Println("This will test sorting performance with different input sizes...")
	fmt.Println()

	summary := t.useCase.RunAllBenchmarks(algorithmName)
	
	for i, result := range summary.Results {
		fmt.Printf("[%d/%d] Generating %s numbers... Sorting... Done in %v ✅\n", 
			i+1, len(summary.Results), pkg.FormatNumber(result.Count), result.Duration)
	}

	pkg.PrintBenchmarkSummary(summary)
}



func (t *Terminal) getMenuChoice(prompt string) string {
	return t.input.ReadString(prompt)
}



func (t *Terminal) askToShowList(result SortResult, count int) {
	if count <= 50 {
		if t.input.ReadYesNo("\nDo you want to see the sorted list? (y/n): ") {
			fmt.Println("\nSorted List:")
			if result.IsArray {
				pkg.PrintSlice(result.SortedArray)
			} else {
				t.printList(result.SortedList)
			}
		}
	} else if count <= 1000 {
		if t.input.ReadYesNo("\nDo you want to see the first 20 numbers of the sorted list? (y/n): ") {
			fmt.Println("\nFirst 20 numbers of sorted list:")
			if result.IsArray {
				pkg.PrintSlicePartial(result.SortedArray, 20)
			} else {
				t.printListPartial(result.SortedList, 20)
			}
		}
	}
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










