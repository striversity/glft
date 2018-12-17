// Section 07 - Review

package main

import (
	"fmt"
)

func main() {
	// send-only channel
	ch := make(chan int, 5)
	ch <- 4
	// populate channel
	producer(ch)
}
func producer(out chan<- int) {
	fmt.Println("Producer")
	fmt.Printf("...ah, go %v from channel \n", <-out)
	for i := 0; i < cap(out); i++ {
		out <- i + 1
	}
}
