package main

import (
	"math"
)

// This file contains several structs/clasess
// The goal is to implement an interface that calculates the area of each shape
// An interface defines the common methods that the classes must implement
// Each class implements the method of the interface with its own code
// Conclusion: The method is the same for all classes but each one implements it different
type Calculation interface {
	calculateArea() float64
	calculatePerimeter() float64
}

type Square struct {
	side float64
}

func (s *Square) calculateArea() float64 {
	return s.side * s.side
}

func (s *Square) calculatePerimeter() float64 {
	return (2 * s.side) + (2 * s.side)
}

type Rectangle struct {
	base   float64
	height float64
}

func (r *Rectangle) calculateArea() float64 {
	return r.base * r.height
}

func (r *Rectangle) calculatePerimeter() float64 {
	return (2 * r.base) + (2 * r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) calculateArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c *Circle) calculatePerimeter() float64 {
	return (2 * math.Pi) * c.radius
}

func calculateShapeArea(calculation Calculation) float64 {
	return calculation.calculateArea()
}

func calculateShapePerimeter(calculation Calculation) float64 {
	return calculation.calculatePerimeter()
}
