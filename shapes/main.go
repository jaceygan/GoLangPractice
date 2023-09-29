package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height, base float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{sideLength: 3.20}
	printArea(s)

	t := triangle{height: 7, base: 4}
	printArea(t)

}
