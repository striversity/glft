// Section 09 - Lecture 02 : Implementing an interface
package main

import "fmt"

type Currency float64

func main() {
	// review adding methods to type
	// ----
	var c Currency = 11.04
	fmt.Println(c.String()) // $11.04
}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}
