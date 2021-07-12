package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rect struct {
	height, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rect) area() float64 {
	return r.height * r.width
}
func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rect{width: 2, height: 2}
	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}
