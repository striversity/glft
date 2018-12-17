// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// incorrect ways of iterating over a channel's values
	// ----
	fillCh(5, 5)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
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
