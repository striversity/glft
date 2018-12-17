// Section 06 - Lecture 04 : Review - critical section
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 40
)

var (
	s            = rand.NewSource(time.Now().Unix())
	r            = rand.New(s)
	wg           sync.WaitGroup
	counter      = 0
	counterMutex sync.Mutex
	m            map[int]int
)

func main() {
	m = make(map[int]int)
	// Preventing section of code from running concurrently
	fmt.Printf("Starting %v goroutines to increment 'counter'\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		updateCounter(i)
	}
	wg.Wait()
	fmt.Printf("Counter = %v\n", counter)
	fmt.Printf("Map m: len = %v, values = %v\n", len(m), m)
}
func updateCounter(id int) {
	// this function is called concurrently to update 'counter'
	wg.Add(1) // need to do this before 'go' keyword
	go func() {
		// prevent concurrent update or sync critical code section
		counterMutex.Lock()
		counter++
		m[id] = m[id] + 1
		counterMutex.Unlock()
		wg.Done()
	}()
}
