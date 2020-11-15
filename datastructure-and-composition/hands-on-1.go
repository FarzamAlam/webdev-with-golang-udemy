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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func calculateArea(s shape) {
	fmt.Println("Area is : ", s.area())
}

func main() {
	c := circle{
		radius: 10,
	}
	calculateArea(c)
	s := square{
		length: 10,
		width:  20,
	}
	calculateArea(s)
}
