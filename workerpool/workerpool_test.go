package workerpool

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func div(count int) (int, error) {
	if count == 0 {
		return 0, fmt.Errorf("can not input zero")
	}
	return (count + 1) / count, nil
}

func sum(arr []int) int {
	res := 0
	for _, i := range arr {
		res += i
	}
	return res
}

func TestWorkerPool(t *testing.T) {
	wp := NewWorkerPool[int](3)
	wp.ErrorCancel = true
	runned := make([]int, 10)
	var task1 Task[int] = func(ct int) error {
		_res, err := div(ct)
		if ct == 5 {
			fmt.Printf("TaskError[%d]:is 5\n", ct)
			return fmt.Errorf("result[%d]: is 5", ct)
		}
		if err != nil {
			fmt.Printf("FuncError[%d]:is %d\n", ct, ct)
			return err
		}
		fmt.Printf("rawResult[%d]:%d\n", ct, _res)
		runned[ct] = 1
		time.Sleep(1 * time.Second)
		return nil
	}
	jobs := make(chan int, 5)
	fmt.Print("Run Start...\n")
	wp.StartTask(task1, jobs)
	assert.Equal(t, 0, sum(runned))
	fmt.Print("Add jobs...\n")
	go func() {
		defer close(jobs)
		for i := 1; i < len(runned); i++ {
			jobs <- i
		}
	}()
	assert.NotEqual(t, 8, sum(runned))
	fmt.Print("wait for complete\n")
	wp.Wait()
	if wp.ErrorCancel {
		assert.NotEqual(t, 8, sum(runned))
	} else {
		assert.Equal(t, 8, sum(runned))
	}
	fmt.Printf("runned:%+v\n", runned)
}
