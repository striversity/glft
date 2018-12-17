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
	// struct types are different if field names are different
	// ----
	var p4 struct {
		Name string
		Age  Age
	}
	fmt.Println(p4)
	// p0 = p4 // error
	// p4 = p0 // error

}
