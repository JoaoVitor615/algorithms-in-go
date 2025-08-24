package pkg

// IsSortedSlice checks if an integer slice is sorted in ascending order
func IsSortedSlice(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	
	return true
}

// IsSortedSliceDesc checks if an integer slice is sorted in descending order
func IsSortedSliceDesc(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	
	return true
}

// IsValidRange checks if a number is within a specified range
func IsValidRange(num, min, max int) bool {
	return num >= min && num <= max
}

// IsEmpty checks if a slice is empty
func IsEmpty(arr []int) bool {
	return len(arr) == 0
}

// HasDuplicates checks if a slice contains duplicate values
func HasDuplicates(arr []int) bool {
	seen := make(map[int]bool)
	
	for _, value := range arr {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	
	return false
}

// CountDuplicates returns the number of duplicate values in a slice
func CountDuplicates(arr []int) int {
	seen := make(map[int]int)
	duplicates := 0
	
	for _, value := range arr {
		seen[value]++
		if seen[value] == 2 { // Count only when we first see a duplicate
			duplicates++
		}
	}
	
	return duplicates
}

// FindMinMax returns the minimum and maximum values in a slice
func FindMinMax(arr []int) (min, max int, valid bool) {
	if len(arr) == 0 {
		return 0, 0, false
	}
	
	min, max = arr[0], arr[0]
	
	for _, value := range arr[1:] {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	return min, max, true
}
