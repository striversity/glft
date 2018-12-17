/* TODO 1 - Given a list of integer values between 0 and 10,
inclusive. Create a text chart using appropriate number
of '*' or '+' to represent each value.

REQUIREMENTS:
1) Your solution *MUST* use at least one function that you
have written.
2) Print the value followed by the represented number of
'*'s or '+'s.

Let's call this an horizontal text chart

NOTE: The functions getListCount() and getListValue() are
provided.
HELP:
	getListCount() returns an uint, which is the total number of values
	getListValue(uint) returns the value of the i-th element in the list
		- IMPORTANT: Each values is betten 0 - 10 inclusively
*/
package main

import "fmt"

func main() {
	printWith := '*'
	for i := 0; i < getListCount(); i++ {
		printDataLine(printWith, getListValue(i))
	}
}

func printDataLine(c rune, v int) {
	fmt.Printf("%4d\t|", v)
	for i := 0; i < v; i++ {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// --- Provided for your assistance
var values = []int{1, 5, 3, 7, 10, 2, 6, 8, 0, 4, 9}

/* getListCount() returns an uint, which is the total
number of values */
func getListCount() int {
	return len(values)
}

/* getListValue(uint) returns the value of the i-th element
in the list
IMPORTANT: Each values is betten 0 - 10 inclusively */
func getListValue(i int) int {
	return values[i]
}
