package main

import "fmt"

func examp1() {
	q := make(chan struct{})
	go func() {
		for {
			select {
			case <-q:
				return
			default:
				//
			}
		}
	}()
	//
	close(q)
}

func examp2() {
	done := make(chan struct{})
	go func(done chan<- struct{}) {
		//
		done <- struct{}{} // перед завершением сообщаем об этом
	}(done)
	<-done // ожидание завершения горутины
}

func examp3() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	num := examp3()
	fmt.Println(<-num)
	fmt.Println(<-num)
	close(num)
}
