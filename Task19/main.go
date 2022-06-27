package main

import "fmt"

func reversStr(str string) string {
	r := []rune(str)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

func main() {
	str1 := "главрыба"
	fmt.Println(reversStr(str1))
}
