package main

import (
	"fmt"
	"strings"
)

func gnirts(str string) string {
	sl := strings.SplitAfter(str, " ")
	rts := ""
	for i := len(sl) - 1; i >= 0; i-- {
		rts += strings.ReplaceAll(sl[i], " ", "") + " "
	}
	return rts
}

func gnirts1(str string) string {
	sl := strings.Split(str, " ")
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return strings.Join(sl, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(gnirts(str))
	fmt.Println(gnirts1(str))
}
