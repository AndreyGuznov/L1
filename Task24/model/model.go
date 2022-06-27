package model

import "example.com/m/task24+/pkg"

type AnotherPoint struct {
	pkg.X
	pkg.Y
}

func NewAnotherPoit(someX pkg.X, someY pkg.Y) *AnotherPoint {
	return &AnotherPoint{
		X: someX,
		Y: someY,
	}
}
