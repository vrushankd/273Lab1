package main

import (
	"fmt"
	"math"
)

type myCircle struct {
	x, y, r float64
}

type myRectangle struct {
	x1, y1, x2, y2 float64
}

type myShape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *myCircle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *myCircle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *myRectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *myRectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l * 2) + (w * 2)
}

func totalArea(shapes ...myShape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...myShape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.perimeter()
	}
	return area
}

func main() {
	c := myCircle{1, 1, 5}
	fmt.Printf("Area of circle is: %.2f \n", c.area())
	fmt.Printf("Perimeter of a circle is: %.2f \n", c.perimeter())
	r := myRectangle{1, 1, 2, 2}
	fmt.Printf("Area of a rectangle is: %.2f \n", r.area())
	fmt.Printf("Perimeter of a rectangle is: %.2f \n", r.perimeter())
	fmt.Printf("Total area of circle and rectangle is: %.2f \n", totalArea(&c, &r))
	fmt.Printf("Total perimeter of circle and rectangle is: %.2f \n", totalPerimeter(&c, &r))
}
