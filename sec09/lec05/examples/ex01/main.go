// Section 09 - Lecture 05 : Type Assertion
package main

import "fmt"

type (
	Person struct {
		name string
	}
)

func main() {
	// static type of a value
	// ----
	var foo interface{}
	var f = 1971.07
	foo = f
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	foo = &Person{name: "Jane Doe"}
	fmt.Printf("foo's value: %v, type: %T\n", foo, foo)
	var goo interface{}
	goo = foo
	fmt.Printf("goo's value: %v, type: %T\n", goo, goo)
	foo = f
	// foo = foo + 10.09   // error -- foo's static type is not float64
	// var g float64 = foo // error -- foo's static type is not float64
}
