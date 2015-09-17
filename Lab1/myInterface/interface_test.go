package main

import (
	"math"
	"testing"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func TestFactorial(t *testing.T) {
	c := myCircle{1, 1, 5}
	if toFixed(c.area(), 2) != 78.54 {
		t.Errorf("Area of circle is incorrect: %.2f \n", c.area())
	}
	if toFixed(c.perimeter(), 2) != 31.42 {
		t.Errorf("Perimeter of a circle is incorrect: %.2f \n", c.perimeter())
	}
	r := myRectangle{1, 1, 2, 2}
	if toFixed(r.area(), 2) != 1.00 {
		t.Errorf("Area of rectangle is incorrect: %.2f \n", r.area())
	}
	if toFixed(r.perimeter(), 2) != 4.00 {
		t.Errorf("Perimeter of a rectangle is: %.2f \n", r.perimeter())
	}
	if toFixed(totalArea(&c, &r), 2) != 79.54 {
		t.Errorf("Total area of circle and rectangle is incorrect: %.2f \n", totalArea(&c, &r))
	}
	if toFixed(totalPerimeter(&c, &r), 2) != 35.42 {
		t.Errorf("Total perimeter of circle and rectangle is incorrect: %.2f \n", totalPerimeter(&c, &r))
	}
}
