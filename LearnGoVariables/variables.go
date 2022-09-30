package main

import "fmt"

func main() {
	var firstName string
	firstName = "Rafael"

	var lastName string = "Hernamperez"

	age := 50

	fmt.Println(firstName, lastName, age)

	var vint int
	var vfloat float64
	var vstring string
	var vbool bool

	fmt.Println(vstring, vint, vfloat, vbool)

	const PI = 3.141592

	fmt.Println("PI = ", PI)
}
