package main

import "fmt"

func sum(v ...int) {
	fmt.Printf("v values are: %v\n", v)
	fmt.Printf("v's type: %T\n", v)
}
func formatter(s string, f ...float64) {
	///
}

var foo = func() {
	fmt.Println("I am the func foo()")
}
var hi = func(s string, c int) {
	for x := 0; x < c; x++ {
		fmt.Printf("Hi, %v\n", s)
	}
}
