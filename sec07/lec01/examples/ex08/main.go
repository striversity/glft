// Section 07 - Lecture 01 : Introduction to Channels
package main

import (
	"fmt"
)

func main() {
	// declaring channels
	// ----
	var ch chan int
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// creating channel with capacity (buffered channels)
	// ----
	ch = make(chan int, 1)
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to a full channel
	// ----
	ch <- 2
	ch <- 3 // failed because the channel is full

}
