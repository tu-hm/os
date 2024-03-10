package semaphore

import "sync"

// Semaphore is a counting semaphore implementation
type Semaphore struct {
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

// NewSemaphore creates a new Semaphore with the given count
func NewSemaphore(count int) *Semaphore {
	s := &Semaphore{
		count: count,
	}
	s.cond = sync.NewCond(&s.mutex)
	return s
}

// Acquire acquires n resources from the semaphore
func (s *Semaphore) Acquire(n int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Wait until enough resources are available
	for s.count < n {
		s.cond.Wait()
	}

	// Acquire the resources
	s.count -= n
}

// Release releases n resources back to the semaphore
func (s *Semaphore) Release(n int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Increase the count and signal any waiting goroutines
	s.count += n
	s.cond.Broadcast()
}
