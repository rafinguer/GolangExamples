package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello, ", name)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	sayHello("Rafael")
	num := sum(10, 22)
	fmt.Println(num)

	fmt.Println("Result: ", sum(num, 3))
}
