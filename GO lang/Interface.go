package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.radius * c.radius }

type Rectangle struct{ width, height float64 }
func (r Rectangle) Area() float64 { return r.width * r.height }

func main() {
	c, r := Circle{}, Rectangle{}

	fmt.Print("Circle radius: ")
	fmt.Scan(&c.radius)

	fmt.Print("Rectangle width and height: ")
	fmt.Scan(&r.width, &r.height)

	PrintArea(c)
	PrintArea(r)
}

func PrintArea(s Shape) { fmt.Printf("Area: %.2f\n", s.Area()) }
