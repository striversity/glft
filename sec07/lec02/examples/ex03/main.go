// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// testing for value sent before closure
	// ----
	sc := make(chan string, 10)
	sc <- "hi"
	sc <- "howdy"
	close(sc)
	var ok bool
	s, ok := <-sc
	fmt.Printf("1st value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("2nd value from sc: %v, ok: %v\n", s, ok)
	s, ok = <-sc
	fmt.Printf("3rd value from sc: %v, ok: %v\n", s, ok)

}
