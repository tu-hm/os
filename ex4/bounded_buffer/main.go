package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	boundedBuffer := NewBoundedBuffer(20)
	wg := sync.WaitGroup{}
	var counter atomic.Int64
	produce := func(product int) {
		for {
			counter.Add(1)
			time.Sleep(time.Duration(rand.Int()%1000) * time.Millisecond)
			boundedBuffer.Produce(int(counter.Load()))
			if counter.Load() > 20 {
				fmt.Println("Done")
				os.Exit(0)
			}
		}
	}
	consume := func(consumer int) {
		for {
			time.Sleep(time.Duration(rand.Int()%1000) * time.Millisecond)
			fmt.Println("Consumer ", consumer, " consume: ", boundedBuffer.Consume())
		}
	}

	for i := 0; i < 1; i++ {
		index := i
		wg.Add(1)
		go consume(index)
	}

	for i := 0; i < 1; i++ {
		index := i
		wg.Add(1)
		go produce(index)
	}

	wg.Wait()
}
