package main

import (
	"fmt"
	"time"
)

func count(ch chan int) {
	i := 0
	for {
		i++
		ch <- i * i
	}
}

func main() {
	t := time.NewTimer(1 * time.Second)
	in := make(chan int)
	go count(in)
	go func() {
		for val := range in {
			fmt.Println(val)
		}
	}()
	<-t.C
}

// func main() {
// 	in := make(chan int, 1)
// 	end := time.After(1 * time.Second)
// 	i := 0
// 	done := make(chan bool, 1)
// 	go func() {

// 		for {
// 			i++
// 			in <- i
// 			select {
// 			case <-end:
// 				done <- true
// 				return
// 			default:
// 				fmt.Println(<-in)
// 			}
// 		}
// 	}()
// 	<-done
// }
