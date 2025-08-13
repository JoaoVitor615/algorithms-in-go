package main

import (
	"fmt"

	"github.com/JoaoVitor615/algorithms-in-go/sorting"
)

func main() {
	fmt.Println("Welcome to the Algorithms-in-Go Terminal! 🚀")
	fmt.Println("Please choose a category to execute:")
	fmt.Println("1. Sorting Algorithms")
	fmt.Println("2. Search Algorithms (Coming Soon)")
	fmt.Println("3. Data Structures (Coming Soon)")
	fmt.Println("4. Dynamic Programming (Coming Soon)")

	var choice string
	fmt.Print("\nEnter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please restart the program.")
		return
	}

	switch choice {
	case "1":
		sorting.RunSortingInterface()
	case "2", "3", "4":
		fmt.Println("This category is coming soon! Stay tuned. 🚀")
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
