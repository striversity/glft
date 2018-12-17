// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// closing a channel
	// ----
	sc := make(chan string, 2)
	sc <- "hello"
	close(sc)
	fmt.Printf("sc: %v, len: %v, cap: %v\n", sc, len(sc), cap(sc))
	// sending to a closed channel
	sc <- "world" // failed sending on a closed channel

}
