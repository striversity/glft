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
	var p0 Person
	fmt.Println(p0)
	// structs with different number of fields
	// ----
	var p3 struct {
		name string
		age  Age
		ssn  ID
	}
	fmt.Println(p3)
	// p0 = p3 // error
	// p3 = p0 // error

}
