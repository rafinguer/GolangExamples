package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	slice1 := list[3:7] // 4, 5, 6, 7
	fmt.Println("slice1: ", slice1)

	slice2 := list[:7] // 1, 2, 3, 4, 5, 6, 7
	fmt.Println("slice2: ", slice2)

	slice3 := list[3:] // 4, 5, 6, 7, 8, 9, 10, 11, 12
	fmt.Println("slice3: ", slice3)
}
