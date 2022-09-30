package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Your name is: ", name)

	fmt.Print("Enter your age:")
	var age int
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("You are ", age, " years old")

	fmt.Print("Enter your salary:")
	var salary float64
	_, err = fmt.Scanf("%f", &salary)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("You gain ", salary, " euros")
}
