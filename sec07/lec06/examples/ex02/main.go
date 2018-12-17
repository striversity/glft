// Section 07 - Lecture 06 : Concurrency Patterns
package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// sink - count numbers
	// ----
	done := make(chan bool)
	d := intGen(done)
	counter(d)
	time.Sleep(1 * time.Second)
	done <- true // signal to go routine to exit gracefully
}

// counter counts numbers receieved on a channel
func counter(in chan int) {
	go func() {
		log.Info("Counter - starting work...")
		start := time.Now()
		var count int
		for range in {
			count++
		}
		fmt.Printf("Counter processed %v items in %v\n", count, time.Since(start))
	}()
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
