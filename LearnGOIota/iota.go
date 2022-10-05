package main

import "fmt"

func main() {
	/*	const (
		monday    = 0
		tuesday   = 1
		wednesday = 2
		thursday  = 3
		friday    = 4
		saturday  = 5
		sunday    = 6
	)*/

	const (
		monday    = iota
		tuesday   = iota
		wednesday = iota
		thursday  = iota
		friday    = iota
		saturday  = iota
		sunday    = iota
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}
