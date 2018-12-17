// Section 06 - Lecture 02 : Waiting & Synchronization
package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
)

func main() {
	// non-ideal way to wait for all goroutines to complete

}
