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
	Println(e)
	Println(11.04)
	Println(ID(19721104))
	Println(SSN("019-72-1104"))
	Println(&Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")})
}

// Println is package's main custom print func
func Println(e Empty) {
	fmt.Printf("[main.Println] e's value: %v, type: %T\n", e, e)
}
