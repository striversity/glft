// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// we can get a pointer to a pointer variable too
	// ----
	count := 1737
	pCount := &count
	var ppCount = &pCount
	fmt.Printf("count's value: %v, address: %v\n", count, &count)
	fmt.Printf("pCount's value: %v, address: %v\n", pCount, &pCount)
	fmt.Printf("ppCount's value: %v, address: %v\n", ppCount, &ppCount)

}
