// Section 02 - Lecture 05 : The "if" statement
package main

import (
	"fmt"
)

const isLoggingEnabled = false

func main() {
	// simple 'if' statement
	if isLoggingEnabled {
		fmt.Println("Yes, log this information for me.")
	}

	x := 5
	if x > 10 {
		fmt.Printf("x > 10: %v\n", true)
	} else {
		fmt.Printf("x > 10: %v\n", false)
	}

	// short var declaration in 'if' statement
	if x := 4; x < 5 {
		fmt.Println("'x' is local to 'if', and is different from 'x' in main.")
	}

	fmt.Println("'x' =", x)

	// 'if-if else-...-else' statement
	x = 11
	if x < 5 {
		fmt.Println("x is < 5")
	} else if x > 10 {
		fmt.Println("x is > 10")
	} else {
		fmt.Println("x is either one of these [5, 6, 7, 8, 9, 10")
	}

	if 'A' < 'B' {
		fmt.Println("'A' is < 'B'")
	}

	today := "Wednesday"
	if "Wednesday" == today {
		fmt.Println("Yippie, today is Wenesday")
	}
	if 'a' < 'A' {
		fmt.Println("'a' < 'A'")
	}
	if "aa" < "ab" {
		fmt.Println(`"aa" < "ab"`)
	}
	if "aa" < "a0" {
		fmt.Println(`"aa" < "a0"`)
	}
}
