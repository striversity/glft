// Section 08 - Lecture 01 : What is a pointer?
package main

import "fmt"

func main() {
	// declaring a pointer
	// ----
	var pCount *int
	fmt.Printf("pCount: %v, pCount: %p\n", pCount, pCount)

	// more examples of declaring pointers
	// ----
	var pBool *bool
	var pString *string
	var pInt64 *int64
	var pMap *map[string]int
	var pChan *chan int
	fmt.Printf("pBool: %v, pString: %v, pInt64: %v, pMap: %v, pChan: %v\n",
		pBool, pString, pInt64, pMap, pChan)

}
