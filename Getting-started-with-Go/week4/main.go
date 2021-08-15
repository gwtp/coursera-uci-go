package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var name string
	fmt.Print("Enter your name: ")
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Bad name input")
		os.Exit(1)
	}
	name = scanner.Text()

	var address string
	fmt.Print("Enter your address: ")
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Bad address input")
		os.Exit(1)
	}
	address = scanner.Text()

	person := map[string]string{
		"name":    name,
		"address": address,
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("failed to marshal person into json format")
		os.Exit(1)
	}
	fmt.Printf("%s\n", bytes)
}
