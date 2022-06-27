package main

import "fmt"

func main() {
	m := make(map[int][]float64)
	key := 0
	sl1 := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, val := range sl1 {
		key = int(val) / 10
		m[key] = append(m[key], val)
	}
	for key := range m {
		fmt.Println(m[key])
	}
}
