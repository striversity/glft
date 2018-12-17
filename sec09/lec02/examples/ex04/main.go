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
	// assigning a value to an interface variable
	// ----
	var c Currency = 11.04
	fmt.Println(c.String())
	fmt.Println(c)

	var mainStringer Stringer
	mainStringer = c
	fmt.Printf("mainStringer's value: %v, type: %T\n", mainStringer, mainStringer)

	var fmtStringer fmt.Stringer
	fmtStringer = c
	fmt.Printf("fmtStringer's value: %v, type: %T\n", fmtStringer, fmtStringer)

}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}
