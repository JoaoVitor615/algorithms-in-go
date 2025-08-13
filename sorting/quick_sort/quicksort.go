package quick_sort

// QuickSort sorts an array using the QuickSort algorithm
// Time Complexity: O(n log n) average, O(n²) worst case
// Space Complexity: O(log n) average, O(n) worst case
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)
	
	quickSortHelper(result, 0, len(result)-1)
	return result
}

// quickSortHelper performs the recursive QuickSort on the array slice
func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := partition(arr, low, high)
		
		// Recursively sort elements before and after partition
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

// partition rearranges the array so that elements smaller than pivot
// are on the left, and elements greater than pivot are on the right
func partition(arr []int, low, high int) int {
	// Choose the rightmost element as pivot
	pivot := arr[high]
	
	// Index of smaller element (indicates right position of pivot)
	i := low - 1
	
	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap elements
		}
	}
	
	// Swap the pivot element with the element at i+1
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSortInPlace sorts an array in-place using the QuickSort algorithm
func QuickSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

// QuickSortWithCustomPivot allows choosing different pivot strategies
type PivotStrategy int

const (
	LastElement PivotStrategy = iota
	FirstElement
	MiddleElement
	RandomElement
)

// QuickSortCustom performs QuickSort with custom pivot selection
func QuickSortCustom(arr []int, strategy PivotStrategy) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	
	quickSortCustomHelper(result, 0, len(result)-1, strategy)
	return result
}

func quickSortCustomHelper(arr []int, low, high int, strategy PivotStrategy) {
	if low < high {
		// Choose pivot based on strategy
		pivotIndex := choosePivot(arr, low, high, strategy)
		
		// Move chosen pivot to end for partitioning
		if pivotIndex != high {
			arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
		}
		
		// Partition and recursively sort
		pi := partition(arr, low, high)
		quickSortCustomHelper(arr, low, pi-1, strategy)
		quickSortCustomHelper(arr, pi+1, high, strategy)
	}
}

func choosePivot(arr []int, low, high int, strategy PivotStrategy) int {
	switch strategy {
	case FirstElement:
		return low
	case MiddleElement:
		return (low + high) / 2
	case RandomElement:
		// For simplicity, we'll use middle as "random"
		// In a real implementation, you'd use rand.Intn()
		return (low + high) / 2
	default: // LastElement
		return high
	}
}
