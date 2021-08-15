package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxConcurrentEating = 2
	maxEatingTimes      = 3
	maxPhilosophers     = 5
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	// number allows us to distinguish each philosopher.
	number          int
	leftCS, rightCS *ChopS
}

func (p *Philo) eat(wg *sync.WaitGroup, canEat chan struct{}, done chan bool) {
	defer wg.Done()
	for i := 0; i < maxEatingTimes; i++ {
		canEat <- struct{}{}
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Starting to eat %d\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		done <- true
		fmt.Printf("finishing eating %d\n", p.number)
	}
}

// Host allows how many philosophers can eat concurrently.
func Host(canEat chan struct{}, done chan bool) {
	var eatingMu sync.Mutex
	var eating int

	go func() {
		for range done {
			eatingMu.Lock()
			eating = eating - 1
			eatingMu.Unlock()
		}
	}()

	for {
		// Need to lock reading and writing to the eating variable.
		eatingMu.Lock()
		// Host will block the canEat requests when maxEating is met.
		if eating < maxConcurrentEating {
			<-canEat
			eating++
		}
		eatingMu.Unlock()
	}

}

// shuffleCSIndex shuffles the index of the left and right chopsticks.
func shuffleCSIndex(index []int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(index), func(i, j int) { index[i], index[j] = index[j], index[i] })
	return index[0], index[1]
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, maxPhilosophers)
	for i := 0; i < maxPhilosophers; i++ {
		cs1, cs2 := shuffleCSIndex([]int{i, (i + 1) % 5})
		philos[i] = &Philo{i + 1, CSticks[cs1], CSticks[cs2]}
	}

	// done channel will allow the philosopher to signal when they are done eating.
	done := make(chan bool)
	defer close(done)

	// canEat channel will allow the philosopher to request permission from the host if they can eat.
	canEat := make(chan struct{})
	defer close(canEat)

	go Host(canEat, done)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, canEat, done)
	}
	wg.Wait()
}
