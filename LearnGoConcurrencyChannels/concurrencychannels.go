package main

import "fmt"

// A Channel is a way to communicate values betewwen goroutines
// Shares memory by communicating
// - Creating a channel -- ch := make(chan <type>)
// - Sending a value through a channel -- ch <- <value>
// - Receive a value from a channel -- V := <-ch

func main() {

	// Execute a single channel example
	runSingleChannel()

	// Execute a single buffered channel example
	runSingleBufferedChannel()

	// Execute a channel with multiple values
	runAdvancedChannel()

}

// Single waitAndSay function
func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}

	c <- true
}

// This function adds n times values to the channel
func sayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}

// This function executes a singel channel example
// Using the channel, we communicate a bool value in order
// to process the corresponding code execution
func runSingleChannel() {
	// Declare and create a channel
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("Hello")

	// Sending a signal to c in order to allow waitAndSay to continue
	c <- true

	// We wait to receive another signal on c before we exit
	<-c
}

func runAdvancedChannel() {
	c := make(chan string)
	go sayHelloMultipleTimes(c, 5)

	// using range in order to process each value from the channel
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c

	fmt.Println("Channel closed?: ", !ok, " - value: [", v, "]")

}

// Buffered channels
//    - Can be considered as buffer of stacked channels
//    - Declaration:
//         ch := make(chan <type>, <number_of_channels)
//    - Senders only block when the buffer is full
//    - Receivers block when the buffer is empty
//    - range can be used to receive values from a channel
//    - close can be used to indicate that channel is closed
func runSingleBufferedChannel() {
	ch := make(chan string, 3)

	ch <- "Hello"
	ch <- "World"
	ch <- "!!!"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
