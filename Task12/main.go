package main

import "fmt"

func compare(mass []string) []string {
	m := make(map[string]int)
	result := []string{}
	for _, val := range mass {
		m[val] = 1
	}
	for key := range m {
		result = append(result, key)
	}
	return result
}

func main() {
	mass1 := [...]string{"cat", "cat", "dog", "cat", "tree"}
	res := compare(mass1[:])
	fmt.Println(res)
}
