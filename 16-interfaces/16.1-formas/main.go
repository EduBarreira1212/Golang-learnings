package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func writeArea(f shape) {
	fmt.Printf("The area of the shape is %0.2f\n", f.area())
}

type rectangle struct {
	hight float64
	width float64
}

func (r rectangle) area() float64 {
	return r.hight * r.width
}

type circle struct {
	ray float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.ray, 2)
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)

}
