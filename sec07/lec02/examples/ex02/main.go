// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// receiving from a closed channel
	// ----
	sc := make(chan string, 2)
	sc <- "hello"
	close(sc)
	s := <-sc
	fmt.Printf("sc 1st value: %v, len: %v, cap: %v\n", s, len(sc), cap(sc))
	s = <-sc
	fmt.Printf("sc 2nd value: %v, len: %v, cap: %v\n", s, len(sc), cap(sc))

}
