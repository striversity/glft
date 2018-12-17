// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// every variable has an address
	// ----
	name := "Jane Smith"
	fmt.Printf("&name: %v -> name: %v\n", &name, name)

}
