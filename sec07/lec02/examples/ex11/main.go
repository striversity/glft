// Section 07 - Lecture 02 : Working with closed Channels
package main

import (
	"fmt"
)

func main() {
	// iterating over a channel using the range operator
	// ----
	fillCh(5, 3)
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch { 
		fmt.Println(v)
	}

	fillCh(5, 1)
	close(ch) // close channel to tell 'range' no more data is expected
	for v := range ch { 
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
