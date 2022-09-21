package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Defining a structure that implements RWMutex synchronization
// RWMutex manages and synchronizes read and write operations
type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func main() {
	mc := MapCounter{m: make(map[int]int)}
	go runWriters(&mc, 10) // passing the address of the original map
	go runReaders(&mc, 10)
	go runReaders(&mc, 10)
	time.Sleep(15 * time.Second)
}

// Writers. Passing map counter by reference using pointers
// Writes made into the original map counter
func runWriters(mc *MapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// Readers. Passing map counter by reference usin pointers
// Reads made over the original map counter
func runReaders(mc *MapCounter, n int) {
	for {
		mc.RLock()
		v := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}
