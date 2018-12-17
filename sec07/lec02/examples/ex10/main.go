// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// iterating over a channel using the range operator
	// ----
	fillCh(5, 3)
	for v := range ch { // fails on deadlock
		fmt.Println(v)
	}

}

var (
	ch chan int
)

func fillCh(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}
