package main

import "fmt"

func main() {
	sl1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 7
	if i < len(sl1) {
		copy(sl1[i:], sl1[i+1:])
		sl1 = sl1[:len(sl1)-1]
		fmt.Println(sl1)
	} else {
		fmt.Println("i out of range")
	}
}
