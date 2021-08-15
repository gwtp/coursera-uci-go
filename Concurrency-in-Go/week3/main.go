package main

/* The main binary will sort a comma seperated list of numbers specified with the --numbers flag. The amount of numbers provided should be divisable by 4.

Example:
  ./main -numbers=1,2,3,4,5,6,7,8

*/
import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var numbers = flag.String("numbers", "", "A comma seperated list of numbers to sort. There should be enough numbers to be divisable by 4")

const partitions = 4

func worker(batch []int, out chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%v\n", batch)
	sort.Ints(batch)
	out <- batch
}

func intSlice(s string) ([]int, error) {
	var result []int
	parts := strings.Split(s, ",")
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func main() {
	flag.Parse()
	input, err := intSlice(*numbers)
	if err != nil {
		fmt.Printf("Please specify comma seperated numbers only: %v", err)
		os.Exit(1)
	}

	inputLen := len(input)
	batchSize := inputLen / partitions
	if (inputLen % partitions) != 0 {
		fmt.Printf("Wrong amount of numbers specified with -numbers. %d not divisable by %d\n", inputLen, partitions)
		os.Exit(1)
	}
	out := make(chan []int)
	var wg sync.WaitGroup
	for i := 0; i < len(input); i += batchSize {
		wg.Add(1)
		go worker(input[i:i+batchSize], out, &wg)
	}

	go func() {
		// Wait for all the workers to complete.
		wg.Wait()
		// Close receive chan once workers are complete otherwise for loop wont know when to terminate.
		close(out)
	}()

	var result []int
	for o := range out {
		result = append(result, o...)
	}
	sort.Ints(result)
	fmt.Println(result)
}
