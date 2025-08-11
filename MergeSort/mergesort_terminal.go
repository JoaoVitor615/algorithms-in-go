package mergesort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RunTerminal is the main function for the Merge Sort user interface.
// It interacts with the user, builds the list, and calls the sorting algorithm.
func RunTerminal() {
	fmt.Println("\n\n[   Merge Sort   ]")
	fmt.Println("\nEnter numbers one by one and press Enter. To stop and sort, just press Enter on an empty line.")

	var head *Node
	var tail *Node

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		newNode := &Node{Value: num}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	if head == nil {
		fmt.Println("The list is empty. Nothing to sort.")
		return
	}

	fmt.Println("\nOriginal List:")
	printList(head)

	// Call the Merge Sort function from the same package.
	sortedList := MergeSort(head)

	fmt.Println("Sorted List:")
	printList(sortedList)
}

// printList is a helper function to print the linked list.
func printList(node *Node) {
	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
