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
	// deferencing ANY pointer, EXCEPT nil value pointer
	// ----
	p := &Student{name: "John Doe", age: 45, ssn: "000-11-2222"}
	pp  := &p
	fmt.Printf("pp's value: %v, type: %T\n", pp, pp)
	fmt.Printf("**pp's value: %v, type: %T\n", **pp, **pp)
}
