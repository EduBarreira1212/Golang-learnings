package shapes

import (
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Ray float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Ray, 2)
}
