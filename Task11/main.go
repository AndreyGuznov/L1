package main

import "fmt"

func compare(a []int, b []int) []int {
	m := make(map[int]int)
	for _, k := range a {
		m[k] += 1
	}
	// fmt.Println(m)
	for _, k := range b {
		// m1[k] += 2
		m[k] += 2
	}
	// fmt.Println(m1)
	result := []int{}

	for key, val := range m {
		if val == 3 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	mass1 := []int{1, 2, 3, 4, 5, 6, 7}
	mass2 := []int{5, 6, 7, 8, 9, 10}
	mass3 := compare(mass1, mass2)
	fmt.Println(mass3)
}
