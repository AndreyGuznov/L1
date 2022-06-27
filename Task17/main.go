package main

import (
	"fmt"
)

func qSort(a []int) []int {
	var i, j, l, r int
	var aI int
	var low, high = make([]int, 0), make([]int, 0)

	low, high = append(low, 0), append(high, len(a)-1)

	for {
		l, low = low[len(low)-1], low[:len(low)-1]
		r, high = high[len(high)-1], high[:len(high)-1]
		// fmt.Println("1)", l, low)
		// fmt.Println("2)", r, high)
		for {
			i, j = l, r
			aI = a[l+(r-l)/2]
			// fmt.Println("aI", aI)
			for {
				for ; a[i] < aI; i++ {
					// fmt.Println("a", a)
				}
				for ; aI < a[j]; j-- {
				}
				if i <= j {
					a[i], a[j] = a[j], a[i]
					i++
					j--
				}
				if i > j {
					break
				}
			}
			if i < r {
				low, high = append(low, i), append(high, r)
			}
			r = j
			if l >= r {
				break
			}
		}
		// fmt.Println(len(low), len(high))
		if len(low)+len(high) == 0 {
			break
		}
	}
	return a
}

func main() {
	mass := []int{1, 10, -12, 5, -10, 2, -2}

	fmt.Println(qSort(mass))
}
