package main

import (
	"fmt"

	"ex4/semaphore"
)

type BoundedBufferInterface interface {
	Produce(data int)
	Consume() int
}

type boundedBufferImpl struct {
	buffer   []int
	capacity int
	mutex    *semaphore.Semaphore
	empty    *semaphore.Semaphore
	full     *semaphore.Semaphore
}

func NewBoundedBuffer(size int) BoundedBufferInterface {
	return &boundedBufferImpl{
		buffer:   make([]int, 0, size),
		capacity: size,
		empty:    semaphore.NewSemaphore(size),
		full:     semaphore.NewSemaphore(0),
		mutex:    semaphore.NewSemaphore(1),
	}
}

func (b *boundedBufferImpl) Produce(data int) {
	b.empty.Acquire(1)
	b.mutex.Acquire(1)
	b.buffer = append(b.buffer, data)
	fmt.Println("Put ", data, " into buffer,", "size is ", len(b.buffer))
	b.mutex.Release(1)
	b.full.Release(1)
}

func (b *boundedBufferImpl) Consume() int {
	b.full.Acquire(1)
	b.mutex.Acquire(1)
	value := b.buffer[0]
	b.buffer = b.buffer[1:]
	fmt.Println("Get ", value, " from buffer,", "size is ", len(b.buffer))
	b.mutex.Release(1)
	b.empty.Release(1)
	return value
}
