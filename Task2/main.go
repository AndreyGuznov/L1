package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func getSquare(i int, mass [5]int) {
	fmt.Println(mass[i] * mass[i])
	wg.Done()
}

func main() {
	m := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < len(m); i++ {
		wg.Add(1)
		go getSquare(i, m)
	}
	wg.Wait()
}
