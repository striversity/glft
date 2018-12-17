// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

func main() {
	// cannot create pointers for consts, numberics, or string literals
	// ----
	const c = 4
	const s = "Hello, world!"
	fmt.Printf("address of const c: %v\n", &c)          // illegal
	fmt.Printf("address of const s: %v\n", &s)          // illegal
	fmt.Printf(`&true: %v\n`, &true)                    // illegal
	fmt.Printf(`&1737: %v\n`, &1737)                    // illegal
	fmt.Printf(`&3.14: %v\n`, &3.14)                    // illegal
	fmt.Printf(`&complex(11,4): %v\n`, &complex(11, 4)) // illegal
	fmt.Printf(`&"hello": %v\n`, &"hello")              // illegal
	type ID string
	fmt.Printf(`&ID("000-11-2222")`, &ID("000-11-2222"))

}
