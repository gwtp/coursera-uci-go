package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	// Check should be case insensitive.
	input = strings.ToLower(input)
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
		os.Exit(0)
	}

	fmt.Println("Not Found!")
	os.Exit(1)
}
