package main

import (
	"fmt"
	"sort"
)

func main() {
	mass1 := []int{4, 4, 5, 5, 6, 7, 8, 9}
	res := sort.SearchInts(mass1, 7) // if not found == len(res)
	fmt.Println(res)
}
