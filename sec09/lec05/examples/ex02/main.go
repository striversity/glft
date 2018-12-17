// Section 09 - Lecture 05 : Type Assertion
package main

import "fmt"

type (
	Person struct {
		name string
	}
)

func main() {
	// using type assertion on for a variable with dynamic type
	// ----
	var foo interface{}
	foo = 1971.07

	var f float64
	// f = float64(foo) -- error
	f = foo.(float64)
	fmt.Printf("f's value: %v, type: %T\n", f, f)

}
