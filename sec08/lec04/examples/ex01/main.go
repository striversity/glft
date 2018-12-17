// Section 08 - Lecture 04 : Pointers and Functions
package main

import "fmt"

func main() {
	// passing pointers to function
	// ----
	count := 1737
	fmt.Printf("count BEFORE incInt(): %v\n", count)
	incInt(count)
	fmt.Printf("count AFTER incInt(): %v\n", count)

}
func incInt(v int) {
	v++ // v = v + 1
}
