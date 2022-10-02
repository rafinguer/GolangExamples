package shapes

// This file contains several structs/clasess
// The goal is to implement an interface that calculates the area of each shape
// An interface defines the common methods that the classes must implement
// Each class implements the method of the interface with its own code
// Conclusion: The method is the same for all classes but each one implements it different
type Calculation interface {
	calculateArea() float64
	calculatePerimeter() float64
}

func CalculateShapeArea(shape Calculation) float64 {
	return shape.calculateArea()
}

func CalculateShapePerimeter(shape Calculation) float64 {
	return shape.calculatePerimeter()
}
