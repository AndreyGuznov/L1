package main

import (
	"fmt"
	"math"

	"example.com/m/task24+/model"
	"example.com/m/task24+/pkg"
)

type Point struct {
	x pkg.X
	y pkg.Y
}

func newPoit(someX pkg.X, someY pkg.Y) *Point {
	return &Point{
		x: someX,
		y: someY,
	}
}

func main() {
	point1 := newPoit(1.5, -5)
	point2 := newPoit(2.5, 8)
	dist := math.Sqrt(math.Pow(float64(point1.x-point2.x), 2.0) + math.Pow(float64(point1.y-point2.y), 2.0))
	fmt.Println(dist)
	point3 := model.NewAnotherPoit(3.0, -5.0)
	point4 := model.NewAnotherPoit(-6.0, 2.0)
	dist1 := math.Sqrt(math.Pow(float64(point3.X-point4.X), 2.0) + math.Pow(float64(point3.Y-point4.Y), 2.0))
	fmt.Println(dist1)
}
