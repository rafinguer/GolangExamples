package main

import (
	"fmt"
	"time"
)

func main() {
	// First method. Execute independently of the second and third methods
	//go tickCounter(1)
	//time.Sleep(5 * time.Second)

	// Second method. Execute independently of the first and third methods
	/*	ticker := time.NewTicker(1 * time.Second)
		go tickCounter2(ticker)
		time.Sleep(5 * time.Second)
		ticker.Stop()
		time.Sleep(10 * time.Second)
		fmt.Println("Exiting...")
	*/
	// Third method. Execute independently of the first and second methods
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go tickCounter3(ticker, done)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting...")
}

func tickCounter(n int) {
	ticker := time.NewTicker(time.Duration(n) * time.Second)
	i := 0
	for t := range ticker.C { // C is the channel of the ticker
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}

func tickCounter2(ticker *time.Ticker) {
	i := 0
	for t := range ticker.C { // C is the channel of the ticker
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}

func tickCounter3(ticker *time.Ticker, done chan bool) {
	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("Count ", i, " at ", t)
		case <-done:
			fmt.Println("done signal")
			break Loop
		}
	}
	fmt.Println("Exiting the tick counter")
}
