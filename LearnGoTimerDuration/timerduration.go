package main

import (
	"fmt"
	"time"
)

func main() {
	//go SlowCounter(2)
	//time.Sleep(15 * time.Second)

	go SlowCounterAfter(2)
	time.Sleep(15 * time.Second)
}

func SlowCounter(n int) {
	i := 0
	// Creating a duration of n seconds
	d := time.Duration(n) * time.Second

	for {
		// Creating a timer for this duration
		t := time.NewTimer(d)
		<-t.C // C is the channel of the timer
		i++
		fmt.Println(i)
	}
}

func SlowCounterAfter(n int) {
	i := 0
	// Creating a duration of n seconds
	d := time.Duration(n) * time.Second

	for {
		// Creating a timer for this duration
		<-time.After(d) // If is launched passed d Duration lapsus
		i++
		fmt.Println(i)
	}
}
