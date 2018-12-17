// Section 07 - Lecture 06 : Concurrency Patterns
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	// fan-out - separating streams
	// ----
	done := make(chan bool)
	numbers := intGen(done)
	s0, s1, s2 := fanOut3(numbers)
	counter(s0)
	counter(s1)
	counter(s2)
	time.Sleep(1000 * time.Millisecond)
	done <- true // signal to go routine to exit gracefully
	wg.Wait()

}

// fanOut3 splits a stream of ints into 3 separate streams
func fanOut3(in chan int) (out0, out1, out2 chan int) {
	wg.Add(1)
	out0 = make(chan int)
	out1 = make(chan int)
	out2 = make(chan int)
	go func() {
		defer wg.Done()
		for v := range in {
			select {
			case out0 <- v:
			case out1 <- v:
			case out2 <- v:
			}
		}
		close(out0)
		close(out1)
		close(out2)
	}()
	return
}

// counter counts numbers receieved on a channel
func counter(in chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
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
	wg.Add(1)
	out = make(chan int)
	go func() {
		defer wg.Done()
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
