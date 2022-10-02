package shapes

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) calculateArea() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (c *Circle) calculatePerimeter() float64 {
	return (2 * math.Pi) * c.Radius
}
