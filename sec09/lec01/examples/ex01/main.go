// Section 09 - Lecture 01 : What is an interface?
package main

import "fmt"

func main() {
	// declaring an interface
	// ----
	type Empty interface {
	}

	// interface variable
	// ----
	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)

}
