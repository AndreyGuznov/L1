package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	res int
)

func getSquare1(i int, sl []int) {
	res += sl[i] * sl[i]
	wg.Done()
}

func main() {
	sl1 := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(sl1); i++ {
		wg.Add(1)
		go getSquare1(i, sl1)
	}
	wg.Wait()
	fmt.Println(res)
}
