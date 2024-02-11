package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLen float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLen: 10}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLen * s.sideLen
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
