package pkg

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// InputReader provides utilities for reading user input
type InputReader struct {
	reader *bufio.Reader
}

// NewInputReader creates a new InputReader instance
func NewInputReader() *InputReader {
	return &InputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ReadString reads a string from user input
func (ir *InputReader) ReadString(prompt string) string {
	if prompt != "" {
		print(prompt)
	}
	input, _ := ir.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// ReadInt reads an integer from user input with validation
func (ir *InputReader) ReadInt(prompt string, min, max int) (int, error) {
	input := ir.ReadString(prompt)
	
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	
	if num < min || num > max {
		return -1, strconv.ErrRange
	}
	
	return num, nil
}

// ReadIntOrDefault reads an integer or returns -1 for invalid input
func (ir *InputReader) ReadIntOrDefault(prompt string, min, max int) int {
	num, err := ir.ReadInt(prompt, min, max)
	if err != nil {
		return -1
	}
	return num
}

// ReadNumbers reads a sequence of numbers until empty line
func (ir *InputReader) ReadNumbers() []int {
	var numbers []int
	
	for {
		input := ir.ReadString("Enter a number: ")
		
		if input == "" {
			break
		}
		
		num, err := strconv.Atoi(input)
		if err != nil {
			println("Invalid input. Please enter an integer.")
			continue
		}
		
		numbers = append(numbers, num)
	}
	
	return numbers
}

// ReadYesNo reads a yes/no response from user
func (ir *InputReader) ReadYesNo(prompt string) bool {
	response := ir.ReadString(prompt)
	return strings.ToLower(response) == "y" || strings.ToLower(response) == "yes"
}
