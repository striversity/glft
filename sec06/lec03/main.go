// Section 06 - Lecture 03 : goroutine pitfalls
package main

import (
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

}
