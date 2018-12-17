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

	// receiving data from a channel
	// ----
	<-ch // blocked receiving from nil channel
}
