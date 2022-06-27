package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 1
	switch val := a.(type) {
	case int:
		fmt.Print("Its int: ", val)
	case string:
		fmt.Print("Its string: ", val)
	case bool:
		fmt.Print("Its bool: ", val)
	case chan struct{}:
		fmt.Print("Its chan: ", val)
	default:
		fmt.Print("I dont know.. ")
	}
	//
	//
	t := reflect.TypeOf(a)
	fmt.Println("\n", t)

	// fmt.Printf("\n%T", a)
}
