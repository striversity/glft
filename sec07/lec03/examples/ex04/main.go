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
	// processing pipeline v2
	// ----
	d := generator()
	d = counter(d)
	d = adder(d, 5)
	consumer(d)
}

// adder takes a channel 'in' and adds the value 'c' to each
// element receieved from the channel, then write the result
// to 'out' channel
func adder(in chan int, c int) (out chan int) {
	out = make(chan int, len(in))
	for v := range in {
		out <- v + c
	}
	close(out)
	return
}

// counter takes a channel 'in' and copies the elemnt to
// 'out' channel, it writes the number of elements copied
func counter(in chan int) (out chan int) {
	out = make(chan int, len(in))
	count := 0
	for v := range in {
		out <- v
		count++
	}
	close(out)
	fmt.Printf("Counted %v elements\n", count)
	return
}

// generator creates a chan on which to reciev random numbers
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
