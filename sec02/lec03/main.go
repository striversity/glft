// Section 02 - Lecture 03 : Strings and rune values
package main

import (
	"fmt"
)

func main() {
	// simple double quoted strings using " and "
	fmt.Println("Hello, World!")
	fmt.Println("Hello, 世界!")

	/* a rune is a 32-bit number which represents a code-point.
	A code-point is how UTF-8 (unicode) characters are encoded.
	A rune is big enough to represent any printable character,
	even in languages like Chinese with thousands of characters. */
	fmt.Println('H')
	fmt.Println('e')
	fmt.Println('l')
	fmt.Println('l')
	fmt.Println('o')
	fmt.Println(',')
	fmt.Println(' ')
	fmt.Println('世')
	fmt.Println('界')
	fmt.Printf("19990 = %c\n", 19990)
	fmt.Printf("19990 = %d\n", 19990)
	fmt.Printf("19990 = %X\n", 19990)
	// fmt.Println('界界') // illegal

	// escape characters
	fmt.Println("Hello,", 10, "World!")
	fmt.Println('\n')
	fmt.Println("Hello, \n World!")
	fmt.Println('\\', "\\")
	fmt.Println('"', "\"")

	// a much longer string that can include just about anything
	fmt.Println(`This is a very long string in golang.
it is called a "raw string" and can span several lines.
NOTE: Strings which are enclosed by "s are not able
to do this, you would have to use an escape charater on the ".`)

	// string concatenation using '+' operator for strings
	fmt.Println("This is another very long string using \"." +
		" But it is getting too lon gfor my screen." +
		"I can add more, but be sure to use the + concatenation" +
		" operator on the same line of the preceeding string.")
}
