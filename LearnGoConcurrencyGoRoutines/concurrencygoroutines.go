package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple way
	go waitAndSay("World")
	fmt.Println("Hello")
	time.Sleep(3 * time.Second)

	// Sharing the value of a variable between threads
	word := "Hello"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Print(word)
	}()

	fmt.Println(word)
	word = "World"
	time.Sleep(3 * time.Second)
}

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}
