// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
	"time"
)

func main() {
	// inserting delays with time.Sleep()
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())

}
