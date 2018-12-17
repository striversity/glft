// Section 07 - Lecture 06 : Concurrency Patterns
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// random number generator
	// ----
	done := make(chan bool)
	d := intGen(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-d)
	}
	done <- true // signal to go routine to exit gracefully

}

// intGen returns a channel on which random ints are produced
func intGen(done chan bool) (out chan int) {
	out = make(chan int)
	go func() {
		for {
			select {
			case <-done:
				close(out)
				return
			case out <- r.Int() % 200:
			}
		}
	}()
	return out
}
