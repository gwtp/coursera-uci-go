package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

const exitKey = "X"

func main() {
	// Before entering the loop, the program should create an empty
	// integer slice of size (length) 3.
	inputs := make([]int, 3)
	for i := 0; true; i++ {
		var input string
		fmt.Printf("Enter an integer or %q to exit: ", exitKey)
		fmt.Scan(&input)
		if input == exitKey {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		// If the following condition is true, we need to grow the
		// underlying array before adding the latest input. We will
		// double it in size.
		if i >= len(inputs) {
			newInputs := make([]int, len(inputs)*2)
			copy(newInputs, inputs)
			inputs = newInputs
		}
		inputs[i] = inputInt

		// The original slice cannot be sorted otherwise we break
		// our ability to add new values by index so we will create
		// a copy, order that and print it.
		printInputs := make([]int, len(inputs))
		copy(printInputs, inputs)
		sort.Ints(printInputs)
		fmt.Print(printInputs)
	}
}
