package quick_sort

import (
	"fmt"
	"reflect"
	"testing"
)

// TestQuickSort runs unit tests for the QuickSort function.
func TestQuickSort(t *testing.T) {
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

			// Call the QuickSort function
			result := QuickSort(tc.input)

			// Verify original input wasn't modified
			if !reflect.DeepEqual(tc.input, originalInput) {
				t.Errorf("QuickSort modified the original input array")
			}

			// Compare the result with the expected output
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("QuickSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestQuickSortInPlace tests the in-place version of QuickSort
func TestQuickSortInPlace(t *testing.T) {
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

			// Call the in-place QuickSort function
			QuickSortInPlace(input)

			// Compare the result with the expected output
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("QuickSortInPlace modified array to %v; want %v", input, tc.expected)
			}
		})
	}
}

// TestQuickSortCustom tests QuickSort with different pivot strategies
func TestQuickSortCustom(t *testing.T) {
	input := []int{4, 2, 5, 1, 3, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	strategies := []PivotStrategy{
		LastElement,
		FirstElement,
		MiddleElement,
		RandomElement,
	}

	strategyNames := []string{
		"LastElement",
		"FirstElement", 
		"MiddleElement",
		"RandomElement",
	}

	for i, strategy := range strategies {
		t.Run(strategyNames[i], func(t *testing.T) {
			result := QuickSortCustom(input, strategy)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("QuickSortCustom with %s strategy: got %v; want %v", 
					strategyNames[i], result, expected)
			}
		})
	}
}

// BenchmarkQuickSort benchmarks the QuickSort function
func BenchmarkQuickSort(b *testing.B) {
	// Create test data
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			// Generate test data
			data := make([]int, size)
			for i := 0; i < size; i++ {
				data[i] = size - i // Reverse sorted for worst case
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Make a copy for each iteration
				testData := make([]int, len(data))
				copy(testData, data)
				b.StartTimer()
				
				QuickSort(testData)
			}
		})
	}
}


