package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func userInput(s *bufio.Scanner) float64 {
	for {
		s.Scan()
		if s.Err() != nil {
			fmt.Print("Bad input, enter value again: ")
			continue
		}
		num, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Print("Bad input, enter value again: ")
			continue
		}
		return num
	}
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 { return (a*(math.Pow(t, 2)*0.5) + (v * t) + s) }
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter acceleration value: ")
	a := userInput(scan)
	fmt.Printf("Enter initial velocity value: ")
	v := userInput(scan)
	fmt.Printf("Enter initial displacement value: ")
	s := userInput(scan)
	fmt.Printf("Enter time value: ")
	t := userInput(scan)

	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}
