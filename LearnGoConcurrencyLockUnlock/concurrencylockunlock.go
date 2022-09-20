package main

import (
	"fmt"
	"sync"
)

// Implementing sync.Mutex for lock and unlock
// Mutex allows you to synchronize the access to
// common values of the threads in a safe way
type safeCounter struct {
	sync.Mutex
	i int
}

func main() {
	sc := new(safeCounter)

	// Creating 100 concurrent threads
	for i := 0; i < 100; i++ {
		go sc.increment()
		go sc.decrement()
	}

	fmt.Println(sc.getValue()) // It must be alway 0
}

// Increment counter between threads in a safe way (blocked mode)
func (sc *safeCounter) increment() {
	sc.Lock()
	sc.i++
	sc.Unlock()
}

// Decrement counter between threads in a safe way (blocked mode)
func (sc *safeCounter) decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func (sc *safeCounter) getValue() int {
	sc.Lock()
	v := sc.i
	sc.Unlock()
	return v
}
