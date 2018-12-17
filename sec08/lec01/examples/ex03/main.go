// Section 08 - Lecture 01 : What is a pointer?
package main

import "fmt"

func main() {
	// declaring pointers to user defined types
	// ----
	type SSN string
	var pSsn *SSN
	fmt.Printf("pSsn: %v\n", pSsn)

	type Person struct {
		fname string
		lname string
		age   uint8
	}
	var pPerson *Person
	fmt.Printf("pPerson: %v\n", pPerson)

}
