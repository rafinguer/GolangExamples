package main

import "fmt"

func main() {
	fmt.Println("--- for min to max incrementing 1 ---")
	// Classical for: from min to max values incrementing iterator
	for i := 1; i <= 10; i++ {
		fmt.Println("Number: ", i)
	}

	// While condition
	fmt.Println("--- while condition ---")

	var i int = 1

	for i <= 10 {
		fmt.Println("Number: ", i)
		i++
	}

	// While condition
	fmt.Println("--- while condition (using continue) ---")

	i = 0

	for i <= 10 {
		i++

		// If noun number skip
		if i%2 != 0 {
			continue
		}

		// Pair number
		fmt.Println("Number: ", i)
	}

	// While true (infinite loop)
	// You must to use break sentence to break
	fmt.Println("--- while true (using break) ---")

	i = 1

	for {
		if i > 10 {
			break
		}

		fmt.Println("Number: ", i)
		i++
	}

}
