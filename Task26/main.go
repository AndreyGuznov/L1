package main

import (
	"fmt"
	"strings"
)

func checkout(str string) bool {
	var m = map[rune]bool{}
	str = strings.ToLower(str)
	for _, w := range str {
		m[w] = true
	}
	return len(m) == len(str)
}

func main() {
	str1 := "abcd"
	fmt.Println(str1, checkout(str1))
	str2 := "abCdefAaf"
	fmt.Println(str2, checkout(str2))
	str3 := "aabcd"
	fmt.Println(str3, checkout(str3))
}
