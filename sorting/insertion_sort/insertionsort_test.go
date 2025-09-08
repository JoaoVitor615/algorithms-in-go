package insertion_sort

import (
	"fmt"
	"reflect"
	"testing"
)

// TestInsertionSort runs unit tests for the InsertionSort function.
func TestInsertionSort(t *testing.T) {
	// Define a series of test cases.
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array with even number of elements",
			input:    []int{4, 2, 5, 1, 3, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Unsorted array with odd number of elements",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicate elements",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
		{
			name:     "Array with all same elements",
			input:    []int{3, 3, 3, 3, 3},
			expected: []int{3, 3, 3, 3, 3},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-5, 2, -3, 8, 1, -1},
			expected: []int{-5, -3, -1, 1, 2, 8},
		},
		{
			name:     "Large random array",
			input:    []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42},
			expected: []int{11, 12, 22, 25, 34, 42, 50, 64, 76, 88, 90},
		},
	}

	// Iterate through each test case.
	for _, tc := range testCases {
		// Run the test in a subtest for clear output.
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of input to ensure original isn't modified
			originalInput := make([]int, len(tc.input))
			copy(originalInput, tc.input)

			// Call the InsertionSort function
			result := InsertionSort(tc.input)

			// Verify original input wasn't modified
			if !reflect.DeepEqual(tc.input, originalInput) {
				t.Errorf("InsertionSort modified the original input array")
			}

			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestInsertionSortOptimized runs unit tests for the InsertionSortOptimized function.
func TestInsertionSortOptimized(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted array (best case)",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array (worst case)",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of input to ensure original isn't modified
			originalInput := make([]int, len(tc.input))
			copy(originalInput, tc.input)

			// Call the InsertionSortOptimized function
			result := InsertionSortOptimized(tc.input)

			// Verify original input wasn't modified
			if !reflect.DeepEqual(tc.input, originalInput) {
				t.Errorf("InsertionSortOptimized modified the original input array")
			}

			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSortOptimized(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestInsertionSortInPlace tests the in-place version of InsertionSort
func TestInsertionSortInPlace(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Unsorted array",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy for in-place sorting
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			// Call the in-place InsertionSort function
			InsertionSortInPlace(input)

			// Compare the result with the expected output
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("InsertionSortInPlace modified array to %v; want %v", input, tc.expected)
			}
		})
	}
}

// TestInsertionSortInPlaceOptimized tests the optimized in-place version of InsertionSort
func TestInsertionSortInPlaceOptimized(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted array (best case)",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy for in-place sorting
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			// Call the optimized in-place InsertionSort function
			InsertionSortInPlaceOptimized(input)

			// Compare the result with the expected output
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("InsertionSortInPlaceOptimized modified array to %v; want %v", input, tc.expected)
			}
		})
	}
}

// TestInsertionSortWithCallback tests InsertionSort with callback functionality
func TestInsertionSortWithCallback(t *testing.T) {
	input := []int{4, 2, 5, 1, 3}
	expected := []int{1, 2, 3, 4, 5}

	insertionCount := 0
	callback := func(arr []int, insertPos, currentPos int) {
		insertionCount++
		// Verify that indices are valid
		if insertPos < 0 || insertPos >= len(arr) || currentPos < 0 || currentPos >= len(arr) {
			t.Errorf("Invalid indices in callback: insertPos=%d, currentPos=%d, len=%d", insertPos, currentPos, len(arr))
		}
	}

	result := InsertionSortWithCallback(input, callback)

	// Verify the result is correct
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InsertionSortWithCallback(%v) = %v; want %v", input, result, expected)
	}

	// Verify that callback was called (insertions occurred)
	if insertionCount == 0 {
		t.Error("Expected callback to be called for insertions, but insertionCount is 0")
	}

	// Test with nil callback (should not panic)
	result2 := InsertionSortWithCallback(input, nil)
	if !reflect.DeepEqual(result2, expected) {
		t.Errorf("InsertionSortWithCallback with nil callback failed")
	}
}

// TestInsertionSortDescending tests the descending order version
func TestInsertionSortDescending(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Unsorted array",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Array with duplicates",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{5, 4, 4, 3, 2, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InsertionSortDescending(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSortDescending(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestInsertionSortWithGap tests the gap-based insertion sort
func TestInsertionSortWithGap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		gap      int
		expected []int
	}{
		{
			name:     "Gap of 1 (normal insertion sort)",
			input:    []int{4, 2, 5, 1, 3},
			gap:      1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Gap of 2",
			input:    []int{5, 2, 4, 6, 1, 3},
			gap:      2,
			expected: []int{1, 2, 4, 3, 5, 6}, // Partially sorted with gap 2
		},
		{
			name:     "Gap of 3",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			gap:      3,
			expected: []int{3, 2, 1, 6, 5, 4, 9, 8, 7}, // Partially sorted with gap 3
		},
		{
			name:     "Invalid gap (0)",
			input:    []int{4, 2, 5, 1, 3},
			gap:      0,
			expected: []int{4, 2, 5, 1, 3}, // Should return original
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InsertionSortWithGap(tc.input, tc.gap)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSortWithGap(%v, %d) = %v; want %v", tc.input, tc.gap, result, tc.expected)
			}
		})
	}
}

// TestInsertionSortStability tests that InsertionSort is stable
func TestInsertionSortStability(t *testing.T) {
	// For testing stability, we need a way to distinguish between equal elements
	// Since we're working with integers, we'll test the behavior with duplicates
	input := []int{3, 1, 3, 2, 3}
	expected := []int{1, 2, 3, 3, 3}

	result := InsertionSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InsertionSort stability test failed: got %v, want %v", result, expected)
	}

	// Test with optimized version
	result2 := InsertionSortOptimized(input)
	if !reflect.DeepEqual(result2, expected) {
		t.Errorf("InsertionSortOptimized stability test failed: got %v, want %v", result2, expected)
	}
}

// BenchmarkInsertionSort benchmarks the InsertionSort function
func BenchmarkInsertionSort(b *testing.B) {
	// Create test data
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			// Generate test data (reverse sorted for worst case)
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = size - i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()

				InsertionSort(testData)
			}
		})
	}
}

// BenchmarkInsertionSortOptimized benchmarks the optimized InsertionSort function
func BenchmarkInsertionSortOptimized(b *testing.B) {
	// Test with already sorted data (best case for insertion sort)
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("sorted_size_%d", size), func(b *testing.B) {
			// Generate sorted test data (best case)
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = i + 1
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()

				InsertionSortOptimized(testData)
			}
		})

		b.Run(fmt.Sprintf("reverse_size_%d", size), func(b *testing.B) {
			// Generate reverse sorted test data (worst case)
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = size - i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()

				InsertionSortOptimized(testData)
			}
		})
	}
}

// BenchmarkInsertionSortInPlace benchmarks the in-place InsertionSort function
func BenchmarkInsertionSortInPlace(b *testing.B) {
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			// Generate test data
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = size - i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()

				InsertionSortInPlace(testData)
			}
		})
	}
}

// BenchmarkInsertionSortBestCase benchmarks InsertionSort on already sorted data
func BenchmarkInsertionSortBestCase(b *testing.B) {
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("best_case_size_%d", size), func(b *testing.B) {
			// Generate already sorted test data (best case)
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = i + 1
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()

				InsertionSort(testData)
			}
		})
	}
}
