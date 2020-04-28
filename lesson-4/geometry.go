package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func SumAreas(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	c1 := &Circle{3}
	c2 := &Circle{5}

	sum := SumAreas(c1, c2, nil)
	fmt.Printf("%.2f\n", sum)
}
