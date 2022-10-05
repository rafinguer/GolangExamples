package main

import (
	"fmt"
	"os"
	"strconv"
)

// Yoy must to pass an argument to this program
// $ go run singleifstatement.go 30
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

	if args := os.Args; len(args) != 2 {
		fmt.Println("Please, give me a number")
	} else {
		fmt.Println("Your number is ", args[1])
	}
}
