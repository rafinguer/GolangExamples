package main

import "fmt"

func main() {
	names := []string{"Rafael", "Edu", "Clemen", "Nerea"}

	// Normal loop, indicating minimum and maximum range
	for i := 0; i < len(names); i++ {
		fmt.Println(i, "-", names[i])
	}

	fmt.Println("---")

	// "For-Each"
	for idx, value := range names {
		fmt.Println(idx, "-", value)
	}
}
