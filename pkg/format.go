package pkg

import (
	"fmt"
	"strings"
)

// FormatNumber formats a number with thousands separators
func FormatNumber(n int) string {
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

// PrintSlice prints an array in bracket format [1, 2, 3]
func PrintSlice(numbers []int) {
	fmt.Print("[")
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println("]")
}

// PrintSlicePartial prints first n elements of an array
func PrintSlicePartial(arr []int, maxCount int) {
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

// PrintHeader prints a formatted header with borders
func PrintHeader(title string) {
	border := strings.Repeat("=", len(title)+4)
	fmt.Println(border)
	fmt.Printf("  %s  \n", title)
	fmt.Println(border)
}

// PrintSubHeader prints a formatted sub-header
func PrintSubHeader(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

// PrintSeparator prints a line separator
func PrintSeparator(length int) {
	fmt.Println(strings.Repeat("-", length))
}
