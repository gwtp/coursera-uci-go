package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	uniqAnimal map[string]Animal
	//userAnimal is a map of user animal name to unique Animal key.
	userAnimal map[string]string
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

func newCow() Animal {
	return &Cow{food: "grass", locomotion: "walk", noise: "moo"}
}

type Cow struct {
	food, locomotion, noise string
}

func (c *Cow) Eat()   { fmt.Println(c.food) }
func (c *Cow) Move()  { fmt.Println(c.locomotion) }
func (c *Cow) Speak() { fmt.Println(c.noise) }

func newBird() Animal {
	return &Bird{food: "worms", locomotion: "fly", noise: "peep"}
}

type Bird struct {
	food, locomotion, noise string
}

func (b *Bird) Eat()   { fmt.Println(b.food) }
func (b *Bird) Move()  { fmt.Println(b.locomotion) }
func (b *Bird) Speak() { fmt.Println(b.noise) }

func newSnake() Animal {
	return &Snake{food: "mice", locomotion: "slither", noise: "hsss"}
}

type Snake struct {
	food, locomotion, noise string
}

func (s *Snake) Eat()   { fmt.Println(s.food) }
func (s *Snake) Move()  { fmt.Println(s.locomotion) }
func (s *Snake) Speak() { fmt.Println(s.noise) }

func init() {
	userAnimal = make(map[string]string)

	uniqAnimal = map[string]Animal{
		"cow":   newCow(),
		"bird":  newBird(),
		"snake": newSnake(),
	}

	return
}

func newAnimal(name, animal string) error {
	// Validate the input.
	if name == "" || animal == "" {
		return errors.New("Bad input, specify `newanimal <name> <animal>`")
	}

	animalLower := strings.ToLower(animal)
	switch animalLower {
	case "bird", "cow", "snake":
	default:
		return fmt.Errorf("Bad input, animal %q is not supported", animal)
	}
	userAnimal[name] = animalLower
	fmt.Println("Created it!")

	return nil
}

func query(name, action string) error {
	// Validate the input.
	if name == "" || action == "" {
		return errors.New("Bad input, specify `query <name> <action>`, try again..")
	}

	ua, ok := userAnimal[strings.ToLower(name)]
	if !ok {
		return fmt.Errorf("Bad input, animal with name %q doesn't exist. Try again..", name)
	}

	a, ok := uniqAnimal[ua]
	if !ok {
		return fmt.Errorf("Oops how embarrasing... This shouldn't happen")
	}

	// Lookup the action.
	switch strings.ToLower(action) {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		return fmt.Errorf("Bad input, action %q doesn't exist. Try again..", action)
	}

	return nil
}

func main() {
	s := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		s, err := s.ReadString('\n')
		if err != nil {
			fmt.Println("Bad input, try again..")
			continue
		}
		if strings.HasSuffix(s, "\n") {
			s = s[:len(s)-1]
		}

		parts := strings.Split(s, " ")
		if len(parts) != 3 {
			fmt.Println("Bad input, please specify: <command> <name> <animal|action>")
		}

		switch strings.ToLower(parts[0]) {
		case "query":
			if err := query(parts[1], parts[2]); err != nil {
				fmt.Println(err)
			}
		case "newanimal":
			if err := newAnimal(parts[1], parts[2]); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Printf("Bad input, command %s does not exist\n", parts[0])
			continue
		}
	}
}
