package insertion_sort

// InsertionSort sorts an array using the Insertion Sort algorithm
// Time Complexity: O(n²) worst and average case, O(n) best case
// Space Complexity: O(1)
func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	// Perform insertion sort
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements that are greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// InsertionSortOptimized sorts an array using an optimized Insertion Sort algorithm
// This version uses binary search to find the insertion position
// Time Complexity: O(n²) worst case (due to shifting), O(n log n) comparisons
// Space Complexity: O(1)
func InsertionSortOptimized(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	// Perform optimized insertion sort
	for i := 1; i < len(result); i++ {
		key := result[i]

		// Find location to insert using binary search
		left, right := 0, i
		for left < right {
			mid := (left + right) / 2
			if result[mid] > key {
				right = mid
			} else {
				left = mid + 1
			}
		}

		// Shift elements to make room for insertion
		for j := i; j > left; j-- {
			result[j] = result[j-1]
		}
		result[left] = key
	}

	return result
}

// InsertionSortInPlace sorts an array in-place using the Insertion Sort algorithm
func InsertionSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Perform insertion sort in-place
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements that are greater than key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// InsertionSortInPlaceOptimized sorts an array in-place using optimized Insertion Sort
func InsertionSortInPlaceOptimized(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Perform optimized insertion sort in-place
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		// Find location to insert using binary search
		left, right := 0, i
		for left < right {
			mid := (left + right) / 2
			if arr[mid] > key {
				right = mid
			} else {
				left = mid + 1
			}
		}

		// Shift elements to make room for insertion
		for j := i; j > left; j-- {
			arr[j] = arr[j-1]
		}
		arr[left] = key
	}
}

// InsertionSortWithCallback sorts an array and calls a callback function after each insertion
// This is useful for visualization or educational purposes
func InsertionSortWithCallback(arr []int, callback func([]int, int, int)) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	// Perform insertion sort with callback
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements that are greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key

		// Call callback function with current state
		if callback != nil {
			callback(result, j+1, i)
		}
	}

	return result
}

// InsertionSortDescending sorts an array in descending order using Insertion Sort
func InsertionSortDescending(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	// Perform insertion sort in descending order
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements that are smaller than key one position ahead
		for j >= 0 && result[j] < key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// InsertionSortWithGap sorts an array using Insertion Sort with a custom gap
// This is used internally by Shell Sort but can be useful on its own
func InsertionSortWithGap(arr []int, gap int) []int {
	if len(arr) <= 1 || gap <= 0 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	// Perform insertion sort with gap
	for i := gap; i < len(result); i++ {
		key := result[i]
		j := i - gap

		// Move elements that are greater than key one gap position ahead
		for j >= 0 && result[j] > key {
			result[j+gap] = result[j]
			j -= gap
		}
		result[j+gap] = key
	}

	return result
}
