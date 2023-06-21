package semaphore

type Semaphore struct {
	capacity chan struct{}
}

func New(maxParallel int) *Semaphore {
	return &Semaphore{
		capacity: make(chan struct{}, maxParallel),
	}
}

func (s *Semaphore) Acquire() {
	s.capacity <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.capacity
}
