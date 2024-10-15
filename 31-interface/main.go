package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (sqr square) area() float64 {
	return sqr.length * sqr.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func info(s shape) float64 {
	return s.area()
}
func main() {

	s1 := square{
		length: 3,
		width:  3,
	}

	c1 := circle{
		radius: 1.5,
	}

	s1area := info(s1)
	c1area := info(c1)

	fmt.Println("Area is", s1area)

	fmt.Println("Area is", c1area)

}
