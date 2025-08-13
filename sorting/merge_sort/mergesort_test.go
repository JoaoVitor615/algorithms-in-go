package merge_sort

import (
	"fmt"
	"reflect"
	"testing"
)

// TestMergeSort runs unit tests for the MergeSort function.
func TestMergeSort(t *testing.T) {
	// Define a series of test cases.
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted list",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted list",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted list with even number of elements",
			input:    []int{4, 2, 5, 1, 3, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Unsorted list with odd number of elements",
			input:    []int{4, 2, 5, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "List with duplicate elements",
			input:    []int{4, 2, 5, 1, 3, 2, 4},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
	}

	// Iterate through each test case.
	for _, tc := range testCases {
		// Run the test in a subtest for clear output.
		t.Run(tc.name, func(t *testing.T) {
			// Create a linked list from the input slice.
			inputList := createList(tc.input)

			// Call the MergeSort function on the list.
			sortedList := MergeSort(inputList)

			// Create the expected linked list for comparison.
			expectedList := createList(tc.expected)

			// Compare the sorted list with the expected list.
			if !compareLists(sortedList, expectedList) {
				t.Errorf("MergeSort(%v) = %v; want %v", tc.input, listToString(sortedList), listToString(expectedList))
			}
		})
	}
}

// createList is a helper function to build a linked list from a slice of integers.
func createList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{Value: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &Node{Value: vals[i]}
		current = current.Next
	}
	return head
}

// compareLists is a helper function to check if two linked lists are equal.
func compareLists(l1, l2 *Node) bool {
	// Convert lists to slices and compare them.
	// This is a robust way to handle the comparison.
	slice1 := listToSlice(l1)
	slice2 := listToSlice(l2)
	return reflect.DeepEqual(slice1, slice2)
}

// listToSlice is a helper function to convert a linked list to a slice.
func listToSlice(head *Node) []int {
	var vals []int
	current := head
	for current != nil {
		vals = append(vals, current.Value)
		current = current.Next
	}
	return vals
}

// listToString is a helper function to format a linked list for error messages.
func listToString(head *Node) string {
	var s string
	current := head
	for current != nil {
		s += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}
	s += "nil"
	return s
}
