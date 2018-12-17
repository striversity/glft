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
	// assignable if types are exactly the same
	// ----
	var i0 int16 = 1971
	var i1 int16
	i1 = i0
	fmt.Printf("i0: %v, i1: %v\n", i0, i1)

	// named types are always different
	// ----
	id0 := ID("71-11-04")
	sku := SKU("72-07-07")
	fmt.Printf("id0: %v, sku: %v\n", id0, sku)
	// id0 = sku // error

	var jane Person
	var bob Individual
	fmt.Printf("jane: %v, bob: %v\n", jane, bob)
	// bob = jane // error

}
