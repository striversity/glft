// Section 06 - Lecture 01 : Introduction to 'goroutines'
package main

import (
	"fmt"
	"time"
)

func main() {
	// creating goroutine from ananymous founction
	// ----
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Foo() - message # %v\n", i)
		}
		producer(3)
	}()
	// give groutines time to complete work
	time.Sleep(1 * time.Millisecond)
}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer # %v - message # %v\n", id, i)
	}
}
