// Section 09 - Lecture 04 : interfaces and functions
package main

import "fmt"

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty interface{}
)

func main() {
	// interface variable as function parameter
	// ----
	var e Empty
	printInfo(e)
	printInfo(11.04)
	printInfo(ID(19721104))
	printInfo(SSN("019-72-1104"))
	printInfo(&Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")})
}

func printInfo(e Empty) {
	fmt.Printf("e's value: %v, type: %T\n", e, e)
}
