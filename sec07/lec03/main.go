// Section 07 - Lecture 03 : Functions and Channels

package main

import (
	"math/rand"
	"time"
)

const (
	chCap = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	// passing channels to functions

}
