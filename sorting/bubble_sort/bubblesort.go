package bubble_sort

// BubbleSort sorts an array using the Bubble Sort algorithm
// Time Complexity: O(n²) worst and average case, O(n) best case (optimized version)
// Space Complexity: O(1)
// (Currently not used in the project)
func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)

	// Perform bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				// Swap elements
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// BubbleSortOptimized sorts an array using an optimized Bubble Sort algorithm
// This version stops early if no swaps are made in a pass (array is already sorted)
// Time Complexity: O(n²) worst case, O(n) best case when array is already sorted
// Space Complexity: O(1)
func BubbleSortOptimized(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)

	// Perform optimized bubble sort
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				// Swap elements
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}

		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}

	return result
}

// BubbleSortInPlace sorts an array in-place using the Bubble Sort algorithm
func BubbleSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}

	n := len(arr)

	// Perform bubble sort in-place
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// BubbleSortInPlaceOptimized sorts an array in-place using optimized Bubble Sort
func BubbleSortInPlaceOptimized(arr []int) {
	if len(arr) <= 1 {
		return
	}

	n := len(arr)

	// Perform optimized bubble sort in-place
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}
}

// BubbleSortWithCallback sorts an array and calls a callback function after each swap
// This is useful for visualization or educational purposes
func BubbleSortWithCallback(arr []int, callback func([]int, int, int)) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	n := len(result)

	// Perform bubble sort with callback
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				// Swap elements
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true

				// Call callback function with current state
				if callback != nil {
					callback(result, j, j+1)
				}
			}
		}

		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}

	return result
}
