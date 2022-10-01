package main

import (
	"fmt"
	"strings"
)

// Single Functions with a single argument
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// Single function with two arguments and returning a single value
func sum(a, b int) int {
	return a + b
}

// Multiple values as argument and returning a single value
func sumMany(nums ...int) int {
	var result int = 0
	for _, n := range nums {
		result += n
	}

	return result
}

// Returning multiple values
func calculate(nums ...int) (int, int, int) {
	//var isum, isub, imul int = 0,0,0
	isum, isub, imul := 0, 0, 1
	for _, n := range nums {
		isum += n
		isub -= n
		imul *= n
	}

	return isum, isub, imul
}

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	f := n * factorial(n-1)

	return f
}

// A closure is a function than returns another function
func repeat(times int) func(text string) string {
	return func(text string) string {
		return strings.Repeat(text, times)
	}
}

// Working with pointers in functions
// func modifyText(text string, newText string) {
func modifyText(text *string, newText string) {
	//fmt.Printf("Memory address: %p\n", &text)
	fmt.Printf("Memory address: %p\n", text)
	//text = newText
	*text = newText
}

func main() {
	sayHello("Rafael")
	num := sum(10, 22)
	fmt.Println(num)

	fmt.Println("sum: ", sum(num, 3))

	fmt.Println("Sum many: ", sumMany(1, 2, 3, 4, 5, 6, 7, 8, 9))

	isum, isub, imul := calculate(1, 2, 3, 4, 5, 6)
	fmt.Println("isum: ", isum, " - isub: ", isub, " - imul: ", imul)

	fmt.Println("factorial(5) = ", factorial(5))

	// Anonymous function
	func() {
		fmt.Println("Hi, I'm an anonymous function")
	}() // This parenthesis executes the anonymous function

	// Anonymous function
	anonymous := func(name string) string {
		result := fmt.Sprintf("Hi, '%s', from an anonymous function", name)
		return result
	}

	fmt.Println(anonymous("Rafael"))

	// Closure function
	text := repeat(5)        // Call to the parent function in the closure
	fmt.Println(text("Go"))  // Call to the child function in the closure
	fmt.Println(text("Foo")) // We can reuse the call to the child function ('text' is a function)

	fmt.Println(repeat(3)("Golang")) // Call to closure function in one step

	// Using pointers with functions
	name := "Rafael"
	newName := "Nerea"

	fmt.Printf("Memory address of %s: %p\n", name, &name)
	fmt.Printf("Name before call: %s\n", name)

	//modifyText(name, newName)
	modifyText(&name, newName)

	fmt.Printf("Name after call: %s\n", name)

}
