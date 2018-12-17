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
	// sync.WaitGroup pitfalls - incorrect initialization
	// ----
	wg.Add(2)
	go producer2(1)
	wg.Wait()

}

// producer2 and main share responsibility of wg
func producer2(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer2 # %v ran for %v\n", id, d)
	wg.Done()
}

func producer(id int) {
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
