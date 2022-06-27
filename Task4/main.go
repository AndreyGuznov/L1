package main

import (
	"context"
	"io"
	"os"
	"strconv"
	"sync"
)

var (
	tasks = make(chan int)
	res   = make(chan int)
)

type PoolOfWorkers struct {
	NumberOf int
}

func New(numberOf int) PoolOfWorkers {
	return PoolOfWorkers{
		NumberOf: numberOf,
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, someTasks <-chan int, someRes chan<- int) {
	defer wg.Done()
	for {
		select {
		case val := <-tasks:
			res <- val
		case _, ok := <-tasks:
			if !ok {
				return
			}
		}
	}
}

func (p *PoolOfWorkers) Work(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < p.NumberOf; i++ {
		wg.Add(1)
		go worker(ctx, wg, tasks, res)
	}
	wg.Wait()
	close(res)
}

func main() {
	var wg sync.WaitGroup
	i := 0
	n := New(1)
	go func() {
		for {
			i++
			tasks <- i
		}
	}()
	go n.Work(context.Background(), &wg)
	for val := range res {
		io.WriteString(os.Stdout, strconv.Itoa(val)+" ")
	}

}
