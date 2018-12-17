// Section 06 - Lecture 04 : Review - critical section
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10000
)

var (
	s       = rand.NewSource(time.Now().Unix())
	r       = rand.New(s)
	wg      sync.WaitGroup
	counter = 0
	// counterMutex sync.Mutex
)

func main() {
	// Preventing section of code from running concurrently
	fmt.Printf("Starting %v goroutines to increment 'counter'\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		updateCounter()
	}
	wg.Wait()
	fmt.Printf("Counter = %v\n", counter)
}
func updateCounter() {
	// this function is called concurrently to update 'counter'
	wg.Add(1) // need to do this before 'go' keyword
	go func() {
		// prevent concurrent update or sync critical code section
		// counterMutex.Lock()
		counter++
		// counterMutex.Unlock()
		wg.Done()
	}()
}
