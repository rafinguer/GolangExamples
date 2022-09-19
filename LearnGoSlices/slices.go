package main

import "fmt"

// Reference documentation:
// https://blog.golang.org/go-slices-usage-and-internals

func main() {

	// Declaring slice without memory allocation
	// var slice []int

	// Declaring slice with memory allocation without values
	// slice := []int{}

	// Declaring slice with memory allocation and reserve length a capacity without values
	// slice := make([]int, 5, 10)
	// len(slice)   => 5
	// cap(slice)   => 10

	// Declaring slice with values
	// slice := []int{1, 2, 3}

	// Creating slice 'mySlice' with preinitialized values
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println("Lenght of mySlice:", len(mySlice))   // 12
	fmt.Println("Capacity of mySlice:", cap(mySlice)) // 12

	// Accessing to a slice:
	//    slice[lowIndex:highIndex]  // from low to high-1
	//    slice[lowIndex:]           // from low to end
	//    slice[:highIndex]          // from beginning to high-1
	slice1 := mySlice[3:7] // 4, 5, 6, 7
	fmt.Println("slice1: ", slice1)

	slice2 := mySlice[:7] // 1, 2, 3, 4, 5, 6, 7
	fmt.Println("slice2: ", slice2)

	slice3 := mySlice[3:] // 4, 5, 6, 7, 8, 9, 10, 11, 12
	fmt.Println("slice3: ", slice3)

	// Append items
	slice1 = append(slice1, 20, 30) // Adding at the end => // 4, 5, 6, 7, 20, 30
	fmt.Println("slice1: ", slice1)

	// Removing items
	//    a = append(a[:i], a[j:]...)

	// Removing one item
	//    a = append(a[:i], a[i+1:]...)

}
