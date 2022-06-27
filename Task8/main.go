package main

import (
	"fmt"
	"log"
	"strconv"
)

func converter(number string) (string, error) {
	bin, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(bin, 2), nil
}

func main() {
	num := ""
	j := 0
	fmt.Printf("Enter the number ")
	fmt.Scanln(&num)
	res, err := converter(num)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Enter number i, 0<i<=%d  ", len(res))
	fmt.Scanln(&j)

	// fmt.Println(res)
	resB := []byte(res)
	if resB[len(string(resB))-j] == 49 {
		resB[len(string(resB))-j] = 48
	} else {
		resB[len(string(resB))-j] = 49
	}
	// fmt.Println(string(resB))

}
