// Section 08 - Lecture 01 : What is a pointer?
package main

import "fmt"

func main() {
	// function pointers are possible too
	// ----
	var pFunc func(int)
	fmt.Printf("pFunc: %p\n", pFunc)

}
