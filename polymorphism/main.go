package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type triangle struct {
	a float64
	h float64
}

type square struct {
	a float64
}

type rectangle struct {
	a float64
	b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (t triangle) area() float64 {
	return 0.5 * t.a * t.h
}

func (s square) area() float64 {
	return s.a * s.a
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Alan : ", shape.area())
	}
}

func main() {
	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 6}
	printArea(t, s, r)

}
