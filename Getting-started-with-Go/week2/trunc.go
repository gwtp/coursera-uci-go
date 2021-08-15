package main

import (
	"fmt"
	"os"
)

func main() {
	userNum := new(float32)
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(userNum)
	if err != nil {
		fmt.Println("Bad input")
		os.Exit(1)
	}
	fmt.Printf("%d\n", int(*userNum))
}
