package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}
type circle struct {
	radious float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		side: 4,
	}

	c1 := circle{
		radious: 3.5,
	}

	info(s1)
	info(c1)
}
