// Section 01 - Lecture 02 : Boolean and numberic values
package main

import (
	"fmt"
)

func main() {
	// boolean values are either true or false
	fmt.Println(true)
	fmt.Println(false)

	// logical operators: ! = NOT, && = AND, and || = OR
	fmt.Println("!false = ", !false)
	fmt.Println("!true = ", !true)
	fmt.Println("false && true = ", false && true)
	fmt.Println("false || true = ", false || true)
	fmt.Println("5 > 10:", 5 > 10)
	fmt.Println(`"abc" == "xyz":`, "abc" == "xyz")

	/* numberic values can be integers (whole numbers),
	floating point (numbers with a decimal point),
	or complex numbers (real and imaginary parts) */
	// examples of integers
	fmt.Println(6)
	fmt.Println("3 + 1 = ", 3+1)
	fmt.Println("-20: ", -20) // negative number
	fmt.Println("-7 + 10 =", -7+10)
	fmt.Println("-7 - -10 =", -7 - -10)
	fmt.Println("-7-(-10) =", -7-(-10))
	fmt.Println("-7 + -10 =", -7+-10)
	fmt.Println("9 / 15 =", 9/15)

	// examples of floating point numbers
	fmt.Println("9.0 / 15 =", 9.0/15)
	fmt.Println("9 / 15.0 = ", 9/15.0)
	fmt.Println("Pi ~=", 3.141592)

	/* golan has built-in support for complex numbers,
	which are numbers with a real an dimaginary part */
	fmt.Println("11 + 4i =", 11+4i)
	fmt.Println("25i =", 25i)
	// fmt.Println("25 i =", 25 i) // illegal, incorrect
	fmt.Println("real(11+4i) =", real(11 + 4i))
	fmt.Println("imag(11+4i) =", imag(11 + 4i))
}
