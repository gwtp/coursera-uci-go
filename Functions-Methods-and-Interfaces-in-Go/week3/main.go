package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initAnimals() map[string]*Animal {
	return map[string]*Animal{
		"cow":   &Animal{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  &Animal{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": &Animal{food: "mice", locomotion: "slither", noise: "hsss"},
	}
}

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat()   { fmt.Println(a.food) }
func (a *Animal) Move()  { fmt.Println(a.locomotion) }
func (a *Animal) Speak() { fmt.Println(a.noise) }

func main() {
	animals := initAnimals()

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">")
		// Get the animal.
		s.Scan()
		if err := s.Err(); err != nil {
			fmt.Println("Bad input, specify animal followed by action eg: cow speak")
			continue
		}

		// Validate the input only contains animal and action.
		parts := strings.Split(s.Text(), " ")
		if len(parts) != 2 {
			fmt.Println("Bad input, specify animal followed by action eg: cow speak")
			continue
		}

		// Lookup the animal.
		a, ok := animals[strings.ToLower(parts[0])]
		if !ok {
			fmt.Printf("Bad input, animal %q doesn't exist. Try again..\n", parts[0])
			continue
		}

		// Lookup the action.
		switch strings.ToLower(parts[1]) {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			fmt.Printf("Bad input, action %q doesn't exist. Try again..\n", parts[1])
			continue
		}
	}
}
