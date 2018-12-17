// Section 06 - Lecture 04 : Review - 10 goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10000 // tested up to 1000000
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	// How many Goroutines can my program have?
	fmt.Printf("Creating %v goroutines", numWorkers)
	for i := 0; i < numWorkers; i++ {
		producer(i)
	}
	wg.Wait()
}

func producer(id int) {
	// do some work in a goroutine
	wg.Add(1) // need to do this before 'go' keyword
	go func() {
		fmt.Printf("Producer # %v - finished\n", id)
		time.Sleep(time.Duration(r.Int()%5000) * time.Nanosecond)
		wg.Done()
	}()
}
