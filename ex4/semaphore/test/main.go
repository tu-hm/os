package main

import (
	"fmt"
	"sync"

	"ex4/semaphore"
)

func main() {
	semaphore1 := semaphore.NewSemaphore(1)
	semaphore2 := semaphore.NewSemaphore(1)
	semaphore3 := semaphore.NewSemaphore(0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("start func 0")
		for {
			semaphore1.Acquire(1)
			fmt.Println(0)
			semaphore2.Release(1)
			semaphore3.Release(1)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("start func 1")
		semaphore2.Acquire(1)
		semaphore1.Release(1)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("start func 2")
		semaphore3.Acquire(1)
		semaphore1.Release(1)
	}()

	wg.Wait()
}
