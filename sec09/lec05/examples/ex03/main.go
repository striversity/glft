// Section 09 - Lecture 05 : Type Assertion
package main

import "fmt"

type (
	Person struct {
		name string
	}
)

func main() {
	// type assertion pitfall
	// ----
	var foo interface{}
	foo = 1971.07

	var f float64
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)
	foo = 7
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)

}
