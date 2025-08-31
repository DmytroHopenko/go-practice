package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Cicrle struct {
	Radius float64
}

func (c Cicrle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Cicrle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printTotals(shapes []Shape) {
	var totalArea, totalPerimeter float64

	for _, s := range shapes {
		totalArea += s.Area()
		totalPerimeter += s.Perimeter()
	}

	fmt.Printf("Total Area: %.2f\n", totalArea)
	fmt.Printf("Total Perimeter: %.2f\n", totalPerimeter)
}

func main() {
	c1 := Cicrle{Radius: 5}
	r1 := Rectangle{Width: 4, Height: 6}

	shapes := []Shape{c1, r1}

	printTotals(shapes)
}
