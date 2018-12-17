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
	// interface variable as function variadic parameter
	// ----
	var e Empty
	Println(e, 11.04,
		ID(19721104),
		SSN("019-72-1104"),
		&Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")})
}

// Println is package's main custom print func
func Println(e ...Empty) {
	fmt.Println("[main.Println]")
	for i := range e {
		fmt.Printf(" parameter[%v]'s value: %v, type: %T\n", i, e[i], e[i])
	}
}
