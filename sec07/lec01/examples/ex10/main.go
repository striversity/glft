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

	// making a channel with more capacity
	// ----
	ch = make(chan int, 2)
	ch <- 2
	ch <- 3
	fmt.Printf("ch len: %v, cap: %v\n", len(ch), cap(ch))

	// receiving data from an empty channel
	// ----
	fmt.Printf("1st value from ch: %v, len: %v\n", <-ch, len(ch))
	fmt.Printf("2nd value from ch: %v, len: %v\n", <-ch, len(ch))
	// next command will fail, since the channel is empty
	fmt.Printf("3rd value from ch: %v, len: %v\n", <-ch, len(ch))

}
