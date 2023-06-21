package semaphore

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	maxParallel := 3
	sem := New(maxParallel)

	var wg sync.WaitGroup

	var jobCounter int
	var counterLock sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		sem.Acquire()
		go func(jobId int) {
			defer sem.Release()
			defer wg.Done()

			counterLock.Lock()
			jobCounter++
			counterLock.Unlock()

			time.Sleep(500 * time.Millisecond)

			counterLock.Lock()
			jobCounter--
			counterLock.Unlock()
		}(i)
	}

	wg.Wait()

	if jobCounter > maxParallel {
		t.Errorf("Number of running jobs exceeded the maximum allowed: %d", jobCounter)
	}
}
