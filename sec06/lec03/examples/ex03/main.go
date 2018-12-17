// Section 06 - Lecture 03 : goroutine pitfalls
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
	// sync.WaitGroup pitfalls - copying sync.WaitGroup value
	// ----
	launchWorkers(5)
	wg.Wait()
	producer(wg, 8)
	wg.Wait()
	// ...
	time.Sleep(1 * time.Second)
}

// producer takes a copy of 'wg', this is incorrect usage
func producer(wg sync.WaitGroup, id int) {
	wg.Add(1)
	go func() {
		// n is a random int between 1 to 1000 inclusive
		n := (r.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer # %v ran for %v\n", id, d)
		wg.Done()
	}()
}

// launchWorkers creates 'n' goroutines using ananymous function
func launchWorkers(n int) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		id := i
		go func() {
			fmt.Printf("I am worker %v\n", id)
			wg.Done()
		}()
	}
}
