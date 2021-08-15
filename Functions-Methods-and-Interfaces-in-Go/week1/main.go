package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(numbers []int, pos int) {
	numbers[pos], numbers[pos+1] = numbers[pos+1], numbers[pos]
}

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		var swapped bool
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return
}

func main() {
	var numbers []int
	s := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 10; {
		fmt.Printf("Enter integer %d: ", i)
		s.Scan()
		if s.Err() != nil {
			fmt.Println("Bad input, try again..")
			continue
		}
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Enter a integer, try again..")
			continue
		}
		numbers = append(numbers, num)
		i++
	}
	bubbleSort(numbers)
	fmt.Println(numbers)
}
