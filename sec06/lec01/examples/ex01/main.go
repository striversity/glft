// Section 06 - Lecture 01 : Introduction to 'goroutines'
package main

import (
	"fmt"
	"time"
)

func main() {
	// creating goroutine from named function
	// ----
	go producer(1)
	go producer(2)
	// give groutines time to complete work
	time.Sleep(1 * time.Millisecond)
}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer # %v - message # %v\n", id, i)
	}
}
