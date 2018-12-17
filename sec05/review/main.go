// Section 05 - Lecture 03 : Review
package main

import (
	"fmt"

	"github.com/striversity/glft/sec05/review/types"
)

/*
1. Adding Methods to structs
2. structs and Functions
2.1. To copy  or not to copy?
3. Field Visibility
3.1 Private vs Public
*/

func main() {
	c := types.Currency(3.955) //
	fmt.Printf("c: %v\n", c)

	jane := types.NewPerson("Jane", "Doe", 21)
	fmt.Println(jane)
	jane = jane.SetAge(22)
	fmt.Println(jane)
	sam := jane
	sam = sam.SetAge(45)
	sam = sam.SetFirstName("Sam")
	fmt.Println(sam)
}
