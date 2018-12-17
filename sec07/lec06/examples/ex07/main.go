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
	// fan-in - combining streams
	// ----
	done := make(chan bool)
	numbers := intGen(done)
	letters := alphaGen(done)
	alphNum := fanIn2(numbers, letters)
	printer(alphNum)
	time.Sleep(10 * time.Millisecond)
	done <- true // signal to go routine to exit gracefully
	done <- true // signal to go routine to exit gracefully
	wg.Wait()

}

// printer - prints the stream of strings
func printer(in chan string) {
	wg.Add(1)
	go func() {
		for v := range in {
			fmt.Printf("%v,", v)
		}
		fmt.Println()
		wg.Done()
	}()
}

// fanIn2 - takes two channels and produces a single stream of the result
func fanIn2(numbers chan int, letters chan rune) (out chan string) {
	wg.Add(1)
	out = make(chan string)
	go func() {
		defer wg.Done()
		for {
			select {
			case n := <-numbers:
				out <- fmt.Sprintf("%v", n)
			case r := <-letters:
				out <- string(r)
			default:
				close(out)
				return
			}
		}
	}()
	return
}

// alphaGen is a generator that spits out random letters between a-z inclusive
func alphaGen(done chan bool) (out chan rune) {
	wg.Add(1)
	out = make(chan rune)
	go func() {
		defer wg.Done()
		letters := []rune("abcdefghijklmnopqrstuvwxyz")
		for {
			select {
			case out <- letters[r.Int()%len(letters)]:
			case <-done:
				close(out)
				return
			}
		}
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
