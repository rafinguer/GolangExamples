package main

import (
	"example/shapes"
	"fmt"
)

func main() {
	s := shapes.Square{Side: 24.5}
	r := shapes.Rectangle{Base: 14.0, Height: 7.75}
	c := shapes.Circle{Radius: 33.45}

	fmt.Println("Area of Square: ", shapes.CalculateShapeArea(&s))
	fmt.Println("Area of Rectangle: ", shapes.CalculateShapeArea(&r))
	fmt.Println("Area of Circle: ", shapes.CalculateShapeArea(&c))
	fmt.Println("---")
	fmt.Println("Perimeter of Square: ", shapes.CalculateShapePerimeter(&s))
	fmt.Println("Perimeter of Rectangle: ", shapes.CalculateShapePerimeter(&r))
	fmt.Println("Perimeter of Circle: ", shapes.CalculateShapePerimeter(&c))

}
