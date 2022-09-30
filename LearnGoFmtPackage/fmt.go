package main

import "fmt"

func main() {
	name := "Rafael"
	age := 33
	salary := 12345.6789

	fmt.Print(name, age, salary)
	fmt.Print("\n")

	fmt.Println(name, age, salary)

	fmt.Printf("Name: %s - Age: %d, Salary: %.2f\n", name, age, salary)

	message := fmt.Sprintf("Name: %s - Age: %d, Salary: %.2f", name, age, salary)

	fmt.Println("Message: ", message)
}
