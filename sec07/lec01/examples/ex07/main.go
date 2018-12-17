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

	// sending data to buffered channel
	// ----
	ch <- 2
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// receiving data from buffered channel
	// ----
	v := <-ch
	fmt.Printf("1st value from ch: %v, len: %v, cap: %v\n", v, len(ch), cap(ch))

}
