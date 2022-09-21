package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Launching 6 threads
	for i := 0; i <= 5; i++ {
		// Increment the WaitGroup counter
		wg.Add(1)

		// Launch the go routine
		go func(i int) {
			// Decrement the WaitGroup counter when the go routine completes
			defer wg.Done()

			// Do some work
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for thread #", i)
		}(i)
	}

	// Wait for all the go routines to complete
	wg.Wait()
}
