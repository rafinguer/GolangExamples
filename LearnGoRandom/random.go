package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the seed according to the time
	// Other way is using rand.seed(time.Now().UnixNano())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Max number (0..99)
	var max int = 100

	var oneNumber int = r1.Intn(max + 1) // 1..100
	fmt.Println("Random number: ", oneNumber)

	var min int = 1

	var otherNumber float64 = r1.Float64() * float64(max-min)
	fmt.Println("Other random number: ", otherNumber)
}
