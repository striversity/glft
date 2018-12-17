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
	// non-ideal way to wait for all goroutines to complete
	// ----
	start := time.Now()
	go producer(1)
	go producer(2)

	// give groutines time to complete work
	time.Sleep(1 * time.Second)
	elapse := time.Since(start)
	fmt.Printf("Non-idea wait took %v\n", elapse)

}

func producer(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n", id, d)
}
