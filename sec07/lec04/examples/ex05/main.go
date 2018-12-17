// Section 07 - Lecture 04 : Channels and Goroutines
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	s           = rand.NewSource(time.Now().Unix())
	r           = rand.New(s)
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

func main() {
	// non-concurrent data producer and consumer
	// ------
	// d := make(chan string, 5)
	// producer(1, d)
	// consumer(d)

	// concurrent data producer and consumer
	// ------
	// d := make(chan string, 5)
	// go producer(1, d)
	// consumer(d)

	// *0 capacity channel* with concurrent data producer and consumer
	// ------
	// d := make(chan string)
	// go producer(1, d)
	// consumer(d)

	// *timed* concurrent data producer and consumer
	// ------
	// d := make(chan string)
	// go producer(1, d)
	// consumer(d)

	// *multiple* concurrent data producers and *single* consumer
	// ------
	// d := make(chan string)
	// producer(1, d)
	// producer(2, d)
	// producer(3, d)
	// go consumer(d)
	// wgProducers.Wait()
	// close(d) // close after all producers

	// *multiple* concurrent data producers and consumers
	// ------
	d := make(chan string)
	producer(1, d)
	producer(2, d)
	producer(3, d)
	consumer(1, d)
	consumer(2, d)
	wgProducers.Wait()
	close(d) // close after all producers
	wgConsumers.Wait() // wait for consumers
}
func consumer(id int, in chan string) {
	wgConsumers.Add(1)
	
	go func(){
		db := make(map[string]int)
		var fields []string
		for v := range in {
			fields = strings.Split(v, ",")
			db[fields[0]]++
		}
		
		for k, v := range db {
			fmt.Printf("Consumer %v - processed %v items for %v\n", id, v, k)
		}
		wgConsumers.Done()
	}()
}
func producer(id int, out chan string) {
	wgProducers.Add(1)

	// launch goroutine to produce data
	go func() {
		i := 1
		end := time.Now().Add(time.Duration(r.Int()%1000) * time.Millisecond)

		for time.Now().Before(end) {
			out <- fmt.Sprintf("Producer: %v, item: %v", id, i)
			i++
		}

		// -- can't call close(out) anymore
		wgProducers.Done()
	}()
}
