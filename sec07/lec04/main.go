// Section 07 - Lecture 04 : Channels and Goroutines
package main

import (
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// non-concurrent data producer and consumer
}
