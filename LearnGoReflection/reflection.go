package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	text := "Text"

	fmt.Println("Type of 12: ", reflect.TypeOf(12))
	fmt.Println("Type of text: ", reflect.TypeOf(text))
	fmt.Println("text is string?: ", reflect.ValueOf(text).Kind() == reflect.String)

	pi := 3.1415
	v := reflect.ValueOf(pi)
	fmt.Println("---")
	fmt.Println("Type of pi: ", v.Type())
	fmt.Println("Is pi a float64:", v.Kind() == reflect.Float64)
	fmt.Println("Value of p.X: ", v.Float())
	fmt.Println("Interface(): ", reflect.ValueOf(pi).Interface())

	p := Point{X: 3, Y: 5}
	fmt.Println("---")
	fmt.Println("Type of p: ", reflect.TypeOf(p))
	fmt.Println("Basic type of p:", reflect.ValueOf(p).Kind())
	fmt.Println("Value of p.X: ", reflect.ValueOf(p))

	fmt.Println("---")
	printType(32)
	printType(pi)
	printType(text)
	printType(true)
	printType(p)
}

// Function that prints the type
// interface {} assummes any type
func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("Data is string: ", i)
	case int:
		fmt.Println("Data is int: ", i)
	case float64:
		fmt.Println("Data is float64: ", i)
	case bool:
		fmt.Println("Data is bool: ", i)
	default:
		fmt.Println("Data is struct: ", reflect.TypeOf(i), "(", reflect.ValueOf(i).Kind(), ")", i)
	}
}
