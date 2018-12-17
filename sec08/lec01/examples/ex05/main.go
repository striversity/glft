// Section 08 - Lecture 01 : What is a pointer?
package main

import "fmt"

type (
	SSN string

	Person struct {
		fname string
		lname string
		age   uint8
	}
)

func main() {
	// more examples of declaring pointers
	// ----
	var pBool *bool
	var pString *string
	var pInt64 *int64
	var pMap *map[string]int
	var pChan *chan int
	fmt.Printf("pBool: %v, pString: %v, pInt64: %v, pMap: %v, pChan: %v\n",
		pBool, pString, pInt64, pMap, pChan)

	// 'nil', a special value without a type, until it needs one
	// ----
	pBool = nil
	pString = nil
	pInt64 = nil
	pMap = nil
	pChan = nil
	// ...
	fmt.Printf("pBool: %v, pString: %v, pInt64: %v, pMap: %v, pChan: %v\n",
		pBool, pString, pInt64, pMap, pChan)

	// other things that are not pointers, can have 'nil' values to
	var s []int
	var m map[int]*Person
	var c chan SSN
	var f func(int, string, *Person) (int, complex64)
	fmt.Printf("s: %v, m: %v, c: %v, f: %p\n", s, m, c, f)
	s = nil
	m = nil
	c = nil
	f = nil
	fmt.Printf("s: %v, m: %v, c: %v, f: %p\n", s, m, c, f)
}
