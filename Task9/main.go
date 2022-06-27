package main

import (
	"fmt"
)

func input(ch chan int, i int) {
	ch <- i
}

func main() {
	m := [5]int{1, 2, 3, 4, 5}
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	end := make(chan struct{}, 1)
	for i := 0; i < len(m); i++ {
		go input(ch1, m[i])
	}
	for i := 0; i < len(m); i++ {
		select {
		case val := <-ch1:
			ch2 <- val * 2
			fmt.Println(<-ch2)
		case <-end:
			break
		}
	}
	end <- struct{}{}
}
