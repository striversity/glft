// Section 09 - Lecture 01 : What is an interface?
package main

import "fmt"

func main() {
	// storing values in an interface variable
	// ----
	type Empty interface {}
	
	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = 7
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = "Hello, World!"
	fmt.Printf("e's value: %v, type: %T\n", e, e)

}
