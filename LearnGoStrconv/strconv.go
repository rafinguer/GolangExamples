package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "45"
	num2 := "1234.5678"

	var f float64
	var b bool

	i, err := strconv.ParseInt(num1, 0, 32)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("String converted to int: ", i)

	f, err = strconv.ParseFloat(num2, 64)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("String converted to float: ", f)

	b, err = strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("String converted to bool: ", b)
}
