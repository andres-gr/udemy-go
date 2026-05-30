package main

import "fmt"

type shape interface {
	getArea() float64
}

type (
	square struct {
		sideLength float64
	}
	triangle struct {
		base   float64
		height float64
	}
)

func main() {
	sq := square{sideLength: 5}
	tr := triangle{base: 10, height: 5}
	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}

func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}
