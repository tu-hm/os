package main

import (
	"fmt"
	"sync"
	"time"

	"ex4/semaphore"
)

const (
	numPhilosophers = 5
	thinking        = 0
	hungry          = 1
	eating          = 2
)

var wg sync.WaitGroup

var marked = make([]bool, numPhilosophers)

type Philosopher struct {
	id            int
	forks         []*semaphore.Semaphore
	state         int
	numMealsEaten int
}

func NewPhilosopher(id int, forks []*semaphore.Semaphore) *Philosopher {
	return &Philosopher{
		id:            id,
		forks:         forks,
		state:         thinking,
		numMealsEaten: 0,
	}
}

func canBeDone() bool {
	//for i := 0; i < numPhilosophers; i++ {
	//	if !marked[i] {
	//		return false
	//	}
	//}
	return true
}

func (p *Philosopher) eat() {
	leftFork := p.forks[p.id]
	rightFork := p.forks[(p.id+1)%numPhilosophers]

	leftFork.Acquire(1)
	fmt.Printf("Philosopher %d picked up left fork\n", p.id)
	rightFork.Acquire(1)
	fmt.Printf("Philosopher %d picked up right fork and is eating\n", p.id)

	p.state = eating
	time.Sleep(time.Second)
	p.state = thinking
	p.numMealsEaten++
	rightFork.Release(1)
	leftFork.Release(1)
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Second)
	p.state = hungry
}

func (p *Philosopher) run(philosophers []*Philosopher) {
	defer wg.Done()
	for {
		//if canBeDone() {
		//	fmt.Println("All philosophers have eaten at least once!")
		//	os.Exit(0)
		//}
		if p.state == thinking {
			p.think()
		} else if p.state == hungry {
			p.waiter(philosophers)
			p.eat()
			marked[p.id] = true
		}
	}
}

func (p *Philosopher) waiter(philosophers []*Philosopher) {
	leftPhilosopher := philosophers[(p.id+numPhilosophers)%numPhilosophers]
	rightPhilosopher := philosophers[(p.id+1)%numPhilosophers]

	for leftPhilosopher.state == eating || rightPhilosopher.state == eating {
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	forks := make([]*semaphore.Semaphore, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = semaphore.NewSemaphore(1)
	}
	marked = make([]bool, numPhilosophers)
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = NewPhilosopher(i, forks)
		wg.Add(1)
		go philosophers[i].run(philosophers)
	}
	wg.Wait()
}
