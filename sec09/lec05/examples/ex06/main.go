// Section 09 - Lecture 05 : Type Assertion
package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	// type switch
	// ----
	for i := 0; i < 10; i++ {
		myPrintln(getValue())
	}

	// type switch pitfall
	// ---
	// var foo = getValue()
	// fmt.Printf("foo's type: %v\n", x.(type)) // error - .(type) outside of switch statement

}
func myPrintln(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Printf("x's type is 'bool', value: %v\n", x.(bool))
	case int:
		fmt.Printf("x's type is 'int', value: %v\n", x.(int))
	case float64:
		fmt.Printf("x's type is 'float64', value: %v\n", x.(float64))
	case complex128:
		fmt.Printf("x's type is 'complex128', value: %v\n", x.(complex128))
	case string:
		fmt.Printf("x's type is 'string', value: %v\n", x.(string))
	case *Person:
		fmt.Printf("x's type is '*Person', value: %v\n", x.(*Person))
	default:
		fmt.Printf("x's type is '<unkown>', type: %T, value: %v\n", x, x)
	}
}

func getValue() interface{} {
	ch := make(chan interface{}, 1)
	select {
	case ch <- true:
	case ch <- 2010:
	case ch <- 9.15:
	case ch <- 7 + 7i:
	case ch <- "Hello world!":
	case ch <- ID("19520925"):
	case ch <- &Person{name: "Jane Doe"}:
	}

	return <-ch
}
