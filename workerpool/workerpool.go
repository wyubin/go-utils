package workerpool

import (
	"context"
	"runtime"
	"sync"
)

type Task[T any] func(job T) error

// use T as job type
type WorkerPool[T any] struct {
	maxWorkers  int
	wg          sync.WaitGroup
	ErrorCancel bool
	cancelFunc  context.CancelFunc
}

// Start task with job channel
func (s *WorkerPool[T]) StartTask(task Task[T], chanJob chan T) {
	var ctx context.Context
	ctx, s.cancelFunc = context.WithCancel(context.Background())
	// prepare workers
	for w := 0; w < s.maxWorkers; w++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for job := range chanJob {
				// check task be cancel by one worker
				if s.ErrorCancel == true {
					select {
					case <-ctx.Done():
						return
					default:
					}
				}
				err := task(job)
				if err != nil {
					s.cancelFunc()
				}
			}
		}()
	}
}

func (s *WorkerPool[T]) Wait() {
	s.wg.Wait()
	s.cancelFunc()
}

func NewWorkerPool[T any](numWorkers int) *WorkerPool[T] {
	numMax := runtime.GOMAXPROCS(0)
	if numWorkers > 0 {
		numMax = numWorkers
	}
	return &WorkerPool[T]{
		maxWorkers: numMax,
	}
}
