package main

import (
	"fmt"
	"myfirstapp/utility"
)

func main() {

	// Variables
	var v1 int // Declaring a variable
	v1 = 1     // Assigning a value to the variable
	var v2 = 2 // Declaring and initializing a variable
	v3 := 3    // Declaring and initializing a variable via inference

	fmt.Printf("v1, v2, v3 = %d %d %d\n", v1, v2, v3)

	// Using packages fmt and utility
	fmt.Println("Hello Go World!!!")
	utility.SayHello()
	utility.Say("Number", 20)

	// Function value (functional programming)
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))

	// Calling to a function passing a number by value
	var x int32 = 5
	change(x)
	fmt.Println("x = ", x) // prints 5. Value is not changed

	// Calling to a function passing a number by reference
	changeX(&x)            // Passing the memory addres of x variable
	fmt.Println("x = ", x) // prints 5. Value is not changed

}

// Normally, arguments in a function are passed by value
func change(x int32) { // this x is local to function and copy the value from call
	x = 10 // only it is affected inside the function
}

// This function received the memory address. Passing the value by reference
func changeX(x *int32) { // The pointer references the memory address from x at main()
	*x = 10 // Change value directly in the memory address
}
