package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

type square struct {
	L float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (s square) area() float64 {
	return s.L * s.L
}

func main() {
	sq := square{
		L: 5,
	}

	ci := circle{
		r: 5.1,
	}

	info(sq)
	info(ci)
}

func info(s shape) {
	fmt.Println(s.area())
}
