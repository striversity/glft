// Section 07 - Review

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

type (
	WorkUnit chan int
)

func main() {
	// channel of channel - a channel on which 'channel of type T' is sent or received
	queue := generateWork()
	// create workers to do work with WorkUnit
	id := 1
	for wu := range queue {
		worker(id, wu)
		id++
	}
	// all work completed, so we can end program now
}
func generateWork() (out <-chan WorkUnit) {
	ch := make(chan WorkUnit, 5)
	for i := 0; i < cap(ch); i++ {
		ch <- genWorkUnit(r.Int() % 10)
	}
	close(ch)
	out = ch
	return
}

// genWorkUnit create some work for a worker
func genWorkUnit(l int) (out WorkUnit) {
	out = make(WorkUnit, l)
	for i := 0; i < l; i++ {
		out <- r.Int() % 100
	}
	close(out)
	return
}

// worker performs some calculations on a WorkUnit
func worker(id int, wu WorkUnit) {
	fmt.Printf("Starting work unit # %v\n", id)
	var sum, count int
	for v := range wu {
		count++
		sum += v
	}
	fmt.Printf("\tProcess %v values for work unit %v\n", count, id)
	if count > 0 {
		fmt.Printf("\tTotal: %v, avg: %v\n", sum, sum/count)
	}
}
