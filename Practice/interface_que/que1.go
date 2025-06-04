package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	//we can avoid using *, because we are not changing its actual value(i.e Radius), instead using its value to create a new value(A)
	A := c.Radius * c.Radius
	return A
}

type Rectange struct {
	Length  float64
	Breadth float64
}

func (r Rectange) Area() float64 {
	A := r.Length * r.Breadth
	return A
}
func main() {
	circle1 := Circle{10}
	fmt.Println("Area of circle=", circle1.Area())
	rectange1 := Rectange{Length: 20, Breadth: 10}
	fmt.Println("Area of rectange=", rectange1.Area())
}
