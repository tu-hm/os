package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"ex4/semaphore"
)

// Data is a shared resource protected by the Reader-Writer problem
type Data struct {
	value int
	mutex sync.Mutex
}

func reader(id int, data *Data, wrtSem, rdSem *semaphore.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Acquire read permission
		rdSem.Acquire(1)

		// Read data (no lock needed since only readers are allowed)
		fmt.Printf("Reader %d: read value %d\n", id, data.value)

		// Release read permission
		rdSem.Release(1)

		// Simulate some work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func writer(id int, data *Data, wrtSem, rdSem *semaphore.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Acquire write permission
		wrtSem.Acquire(1)

		// Modify data
		data.mutex.Lock()
		data.value++
		fmt.Printf("Writer %d: wrote value %d\n", id, data.value)
		data.mutex.Unlock()

		// Release write permission
		wrtSem.Release(1)

		// Simulate some work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	data := &Data{value: 0}
	wrtSem := semaphore.NewSemaphore(1)
	rdSem := semaphore.NewSemaphore(5)

	var wg sync.WaitGroup

	// Start 3 readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go reader(i, data, wrtSem, rdSem, &wg)
	}

	// Start 2 writers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go writer(i, data, wrtSem, rdSem, &wg)
	}

	// Wait for goroutines to finish (in this case, they will run indefinitely)
	wg.Wait()
}
