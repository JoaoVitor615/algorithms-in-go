package bubble_sort

import (
	"fmt"
	"reflect"
	"testing"
)

// TestBubbleSort runs unit tests for the BubbleSort function.
func TestBubbleSort(t *testing.T) {
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

			// Call the BubbleSort function
			result := BubbleSort(tc.input)

			// Verify original input wasn't modified
			if !reflect.DeepEqual(tc.input, originalInput) {
				t.Errorf("BubbleSort modified the original input array")
			}

			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestBubbleSortOptimized runs unit tests for the BubbleSortOptimized function.
func TestBubbleSortOptimized(t *testing.T) {
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

			// Call the BubbleSortOptimized function
			result := BubbleSortOptimized(tc.input)

			// Verify original input wasn't modified
			if !reflect.DeepEqual(tc.input, originalInput) {
				t.Errorf("BubbleSortOptimized modified the original input array")
			}

			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BubbleSortOptimized(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestBubbleSortInPlace tests the in-place version of BubbleSort
func TestBubbleSortInPlace(t *testing.T) {
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

			// Call the in-place BubbleSort function
			BubbleSortInPlace(input)

			// Compare the result with the expected output
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("BubbleSortInPlace modified array to %v; want %v", input, tc.expected)
			}
		})
	}
}

// TestBubbleSortInPlaceOptimized tests the optimized in-place version of BubbleSort
func TestBubbleSortInPlaceOptimized(t *testing.T) {
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

			// Call the optimized in-place BubbleSort function
			BubbleSortInPlaceOptimized(input)

			// Compare the result with the expected output
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("BubbleSortInPlaceOptimized modified array to %v; want %v", input, tc.expected)
			}
		})
	}
}

// TestBubbleSortWithCallback tests BubbleSort with callback functionality
func TestBubbleSortWithCallback(t *testing.T) {
	input := []int{4, 2, 5, 1, 3}
	expected := []int{1, 2, 3, 4, 5}
	
	swapCount := 0
	callback := func(arr []int, i, j int) {
		swapCount++
		// Verify that indices are valid
		if i < 0 || i >= len(arr) || j < 0 || j >= len(arr) {
			t.Errorf("Invalid indices in callback: i=%d, j=%d, len=%d", i, j, len(arr))
		}
	}

	result := BubbleSortWithCallback(input, callback)

	// Verify the result is correct
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSortWithCallback(%v) = %v; want %v", input, result, expected)
	}

	// Verify that callback was called (swaps occurred)
	if swapCount == 0 {
		t.Error("Expected callback to be called for swaps, but swapCount is 0")
	}

	// Test with nil callback (should not panic)
	result2 := BubbleSortWithCallback(input, nil)
	if !reflect.DeepEqual(result2, expected) {
		t.Errorf("BubbleSortWithCallback with nil callback failed")
	}
}

// TestBubbleSortStability tests that BubbleSort is stable
func TestBubbleSortStability(t *testing.T) {
	// For testing stability, we need a way to distinguish between equal elements
	// Since we're working with integers, we'll test the behavior with duplicates
	input := []int{3, 1, 3, 2, 3}
	expected := []int{1, 2, 3, 3, 3}

	result := BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort stability test failed: got %v, want %v", result, expected)
	}

	// Test with optimized version
	result2 := BubbleSortOptimized(input)
	if !reflect.DeepEqual(result2, expected) {
		t.Errorf("BubbleSortOptimized stability test failed: got %v, want %v", result2, expected)
	}
}

// BenchmarkBubbleSort benchmarks the BubbleSort function
func BenchmarkBubbleSort(b *testing.B) {
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
				
				BubbleSort(testData)
			}
		})
	}
}

// BenchmarkBubbleSortOptimized benchmarks the optimized BubbleSort function
func BenchmarkBubbleSortOptimized(b *testing.B) {
	// Test with already sorted data (best case for optimized version)
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
				
				BubbleSortOptimized(testData)
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
				
				BubbleSortOptimized(testData)
			}
		})
	}
}

// BenchmarkBubbleSortInPlace benchmarks the in-place BubbleSort function
func BenchmarkBubbleSortInPlace(b *testing.B) {
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
				
				BubbleSortInPlace(testData)
			}
		})
	}
}

