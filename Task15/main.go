package main

var justString string

func createHugeString(i int) string {
	sl := make([]rune, i)
	// write some text here to sl
	return string(sl)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	// fmt.Println(justString)
}
