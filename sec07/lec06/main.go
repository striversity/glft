// Section 07 - Lecture 06 : Concurrency Patterns
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
	// random number generator
}

