package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, height float64
}

type Triangle struct {
	base, height, a, b, c float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.length + r.height) * 2
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func Render(s Shape) {
	fmt.Printf("Shape has Area: %v and Perimeter: %v\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{4}
	r := Rectangle{3, 4}
	t := Triangle{3, 4, 2, 2, 3}

	Render(c)
	Render(r)
	Render(t)
}
