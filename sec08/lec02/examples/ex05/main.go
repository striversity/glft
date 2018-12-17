// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// create a pointer using new()
	// ----
	pCount := new(int) // ==> count := 0; pCount := &count
	fmt.Printf("pCount's value: %v, type: %T\n", pCount, pCount)

	pName := new(string) // ==> name := ""; pName := &name
	fmt.Printf("pName's value: %v, type: %T\n", pName, pName)

}
