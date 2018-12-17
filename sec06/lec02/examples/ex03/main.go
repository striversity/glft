// Section 06 - Lecture 02 : Waiting & Synchronization
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	// ideal way to wait for all goroutines to complete
	// ----
	start := time.Now()
	wg.Add(4)
	go producer(1)
	go producer(2)
	go producer(3)
	go producer(4)

	// wait for groutines to complete work
	wg.Wait()
	elapse := time.Since(start)
	fmt.Printf("Idea wait took %v\n", elapse)
	
	// reusing waitgroup variable
	start = time.Now()
	wg.Add(1)
	go producer(5)
	wg.Wait()
	elapse = time.Since(start)
	fmt.Printf("Producer(5) took %v\n", elapse)
}

func producer(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer2 # %v ran for %v\n", id, d)
	wg.Done()
}
