// Section 08 - Lecture 03 : Using Pointers
package main

import "fmt"

type(
	ID      string
	Student struct {
		name string
		age  uint8
		ssn  ID
	}
)
func main() {
	// pointer pitfall, CANNOT deference nil value pointer
	// ----
	var p *Student
	fmt.Printf("p's value: %v, type: %T\n", p, p)
	fmt.Printf("*p's value: %v, type: %T\n", *p, *p)  // causes panic
}
