// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// review -- declaring a pointer
	// ----
	var pCount *int
	fmt.Printf("pCount: %v, pCount: %T\n", pCount, pCount)

	// creating pointer from a variable, using & => address-of operator
	// ----
	var count = 1737
	pCount = getAddressOfVariable(count) // doesn't work, for illustration ONLY
	pCount = &count                      // kind-of like: pCount = getAddressOfVariable(count)
	fmt.Printf("pCount: %v, pCount: %p, pCount(type): %T\n", pCount, pCount, pCount)
}

// getAddressOfVariable implements the 'address-of operator'/&
func getAddressOfVariable(i int) *int {
	// do some magic to get the address of variable passed to 'i' and return it
	return nil
}
