package main

import (
	"fmt"

	mergesort "github.com/JoaoVitor615/algorithms-in-go/MergeSort"
)

func main() {
	fmt.Println("Welcome to the Algorithms-in-Go Terminal! 🚀")
	fmt.Println("Please choose an algorithm to execute:")
	fmt.Println("1. Merge Sort")
	fmt.Println("2. (Other algorithms will go here)")

	var choice string
	fmt.Print("\nEnter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please restart the program.")
		return
	}

	switch choice {
	case "1":
		mergesort.RunTerminal()
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
