// Section 07 - Review

package main

import (
	"fmt"
)

func main() {
	// receive-only channel
	// ----
	ch := make(chan int, 5)
	// give it so a function that only needs to receive it
	printer(ch)
}
func printer(in <-chan int) {
	fmt.Println("Printer")
	in <- 5
	fmt.Println("...ah, added some data to channel before printing")
	close(in)
	for v := range in {
		fmt.Println(v)
	}
}
