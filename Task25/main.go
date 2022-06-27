package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	ch := time.Tick(1 * time.Second)
	for range ch {
		i--
		fmt.Println(i)
		if i == 0 {
			break
		}
	}
}
