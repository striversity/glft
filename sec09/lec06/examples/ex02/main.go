// Section 09 - Lecture 06 : Assignability
package main

import "fmt"

type (
	ID     string
	SKU    string
	Person struct {
		name string
		age  Age
	}
	Individual struct {
		name string
		age  Age
	}
	Age uint8
)

func main() {
	// assignment between named and unamed types
	// ----
	var p0 Person
	var p1 struct {
		name string
		age  Age
	}
	// an unamed type is assignable if it is EXACTLY like the name type
	// ----
	p0 = p1
	p1 = p0

	// unamed type struct are not assignable if the order of fields are not EXACT
	// ----
	var p2 struct {
		age  Age
		name string
	}
	fmt.Println(p2)
	// p0 = p2 // error
	// p2 = p0 // error

}
