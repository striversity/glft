// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
	"time"
)

func main() {
	// get a notification from a channel with time.After()
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	<-time.After(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())

}
