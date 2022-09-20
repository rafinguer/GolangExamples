package main

import "fmt"

func main() {
	printType("text")
	printType(32)
	printType(23.56)
	printType(false)
}

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
	}
}
