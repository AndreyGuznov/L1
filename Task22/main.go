package main

import (
	"fmt"
	"math"
)

func sumDiff(a1, a2, b1, b2 int) (string, string) {

	c := float64(a2 - b2)
	switch {
	case c == 0:
		return fmt.Sprintf("Res of Addition: %d*10^%d", a1+b1, a2), fmt.Sprintf("Res of Addition: %d*10^%d", a1-b1, a2)
	case c > 0:
		return fmt.Sprintf("Res of Addition: %f*10^%d", float64(a1*int(math.Pow(10.0, c))+b1)/math.Pow(10.0, c), a2), fmt.Sprintf("Res of Substraction: %f*10^%d", float64(a1*int(math.Pow(10.0, c))-b1)/math.Pow(10.0, c), a2)
	default:
		c = float64(b2 - a2)
		return fmt.Sprintf("Res of Addition: %f*10^%d", float64(a1*int(math.Pow(10.0, c))+b1)/math.Pow(10.0, c), b2), fmt.Sprintf("Res of Substraction: %f*10^%d", float64(a1*int(math.Pow(10.0, c))-b1)/math.Pow(10.0, c), b2)
	}

}

// -9223372036854775808 to 9223372036854775807 10^18
func main() {
	var a1, a2, b1, b2 int
	fmt.Printf("Enter integer path of first number ")
	fmt.Scanln(&a1)
	fmt.Printf("Enter power of a number with a natural exponent for 10 of first number ")
	fmt.Scanln(&a2)
	fmt.Printf("Enter integer path of second number ")
	fmt.Scanln(&b1)
	fmt.Printf("Enter power of a number with a natural exponent for 10 of second number ")
	fmt.Scanln(&b2)
	resDiv := fmt.Sprintf("Res of Divide: %f*10^%d ", float64(a1)/float64(b1), a2-b2)
	resMult := fmt.Sprintf("Res of Multiply %d*10^%d", a1*b1, a2+b2)
	sum, sub := sumDiff(a1, a2, b1, b2)
	fmt.Println(resDiv)
	fmt.Println(resMult)
	fmt.Println(sum)
	fmt.Println(sub)
}
