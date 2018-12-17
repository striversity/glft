// Section 06 - Lab - timing
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
	// How long does this take?
	fmt.Println("Work started")
	start := time.Now()
	// do some work ...
	time.Sleep(time.Duration(r.Int()%5000) * time.Millisecond)
	end := time.Now()
	fmt.Println("Work completed")
	elapse := end.Sub(start)
	fmt.Printf("Time taken to do work was: %v\n", elapse)
}
