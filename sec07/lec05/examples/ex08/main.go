// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
	"time"
)

func main() {
	// timing out after waiting on a channel
	// ----
	d := producer()
	consumer(d)
}
func consumer(in chan string) {
	for {
		alarm := time.After(2 * time.Millisecond)
		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("No more data from 'in'")
				return
			}
			fmt.Printf("Consumer got - %v", m)
		case <-alarm:
			fmt.Println("Timedout waiting on for data")
			return
		}
	}
}
func producer() (out chan string) {
	out = make(chan string)
	go func() {
		count := 1
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}
		time.Sleep(3 * time.Millisecond)
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}
		close(out)
	}()
	return
}
