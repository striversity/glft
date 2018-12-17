// Section 07 - Lecture 03 : Functions and Channels

package main

import (
	"fmt"
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
	// returning channels from functions
	// ----
	d := generator()
	consumer(d)

}

// generator creates a chan on which caller receives random numbers
func generator() (out chan int) {
	fmt.Println("Generator of random ints")
	out = make(chan int, chCap)
	for i := 0; i < cap(out); i++ {
		out <- r.Int() % 200
	}
	close(out)
	return
}

func consumer(nums chan int) {
	for v := range nums {
		fmt.Printf("Consumer got: %v\n", v)
	}
}

