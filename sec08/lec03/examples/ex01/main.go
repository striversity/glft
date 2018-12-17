// Section 08 - Lecture 03 : Using Pointers
package main

import "fmt"

func main() {
	// review -- creating a pointer
	// ----
	count := 1737
	pCount := &count
	fmt.Printf("count's value: %v, type: %T\n", count, count)
	fmt.Printf("pCount's value: %v, type: %T\n", pCount, pCount)
	
}
