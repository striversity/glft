// Section 07 - Lecture 03 : Functions and Channels

package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/mgutz/logxi/v1"
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
	// ----
	d := make(chan int, chCap)
	if d == nil {
		log.Error("Can't allocate channel")
	}
	producer(d)
	consumer(d)

}

func consumer(nums chan int) {
	for v := range nums {
		fmt.Printf("Consumer got: %v\n", v)
	}
}

// producer sends between 1 to cap(nums) random integers into the channel 'nums'
func producer(nums chan int) {
	n := r.Int()%cap(nums) + 1
	for i := 0; i < n; i++ {
		nums <- r.Int() % 200
	}
	close(nums)
}
