package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Value() int {
	return c.num
}

func (c *Counter) Inc() {
	c.Lock()
	c.num += 1
	c.Unlock()
}

func Incrimentator(count *Counter, theEnd chan struct{}) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(num int, count *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			count.Inc()
			// fmt.Printf("%d ", num)
		}(i, count, &wg)
	}
	wg.Wait()
	theEnd <- struct{}{}
	close(theEnd)
}

func main() {
	theEnd := make(chan struct{})
	count := &Counter{
		num: 0,
	}
	go Incrimentator(count, theEnd)
	select {
	case <-theEnd:
		fmt.Println(count.Value())
	}
}
