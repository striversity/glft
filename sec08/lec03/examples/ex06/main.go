// Section 08 - Lecture 03 : Using Pointers
package main

import "fmt"

type (
	ID      string
	Student struct {
		name string
		age  uint8
		ssn  ID
	}
)

func main() {
	// pointer pitfall, gaurding against nil pointer deference
	// ----
	var p *Student
	fmt.Printf("p's value: %v, type: %T\n", p, p)
	
	if p != nil { // check pointer for nil before using
		fmt.Printf("*p's value: %v, type: %T\n", *p, *p)
		fmt.Printf("Student Name (p.name): %v\n", p.name)
	}
}
