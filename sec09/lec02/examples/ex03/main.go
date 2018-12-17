// Section 09 - Lecture 02 : Implementing an interface
package main

import "fmt"

type (
	Currency float64
	Stringer interface {
		String() string
	}
)

func main() {
	// Currency implements interface main.Stringer
	// ----
	var c Currency = 11.04
	fmt.Println(c.String()) // explicit call to Currency's String() method

	// Currency ALSO implements interface fmt.Stringer
	// ----
	fmt.Println(c) // implicitly call to Currency's String(), call by fmt.PrintX()

}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}
