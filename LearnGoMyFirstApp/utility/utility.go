package utility

import "fmt"

// Public function. Exportable
func SayHello() {
	fmt.Println("Hello, World!!!")
}

// Private function. Non exportable
func sayRafa() {
	fmt.Println("Rafa")
}

func Say(s string, i int) {
	fmt.Printf("%s %d\n", s, i)
}

/*
func addSubtractMultiply(a, b int) (addition, subtraction, multiplication int) {
	// or return a+b, a-b, a*b
	addition = a + b
    subtraction = a - b
    multiplication = a * b
    return
}
*/

func addSubtractMultiply(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
