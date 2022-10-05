package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := "50"
	/*
		n, err := strconv.ParseInt(age, 0, 64)

		if err != nil {
			fmt.Println("There was an error parsing age. Error:", err)
		} else {
			fmt.Println("There was NO error parsing age. Age is ", n)
		}
	*/

	if n, err := strconv.ParseInt(age, 0, 64); err == nil {
		fmt.Println("There was NO error parsing age. Age is ", n)
	}
}
