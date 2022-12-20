package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	circumference() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) circumference() float64 {
	return 2 * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) circumference() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func getCircumference(s Shape) float64 {
	return s.circumference()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle = %f \n", getArea(circle))
	fmt.Printf("Rectangle = %f \n", getArea(rectangle))

	fmt.Printf("Circle = %f \n", getCircumference(circle))
	fmt.Printf("Rectangle = %f \n", getCircumference(rectangle))
}
