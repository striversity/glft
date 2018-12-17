// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// test reading from an empty channel without closing
	// ----
	sc := make(chan string, 4)
	sc <- "hi"
	s, ok := <-sc
	fmt.Printf("1st value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc // will fail
	fmt.Printf("2nd value from sc: %v, ok: %v\n", s, ok)

}
