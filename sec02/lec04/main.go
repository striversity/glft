// Section 01 - Lecture 04 : Constants and variables
package main

import (
	"fmt"
)

const (
	pi                       = 3.14159
	ageLimit           uint8 = 18
	defaultCountry           = "United States of America"
	isDebuggingEnabled       = true
	_favChar                 = 'P'
)

// w:= 20 illegal/incorrect
var (
	sot  = 9
	dir  = "20 deg South"
	name = "VLA"
)

func main() {
	fmt.Println(pi)
	fmt.Println(pi + 9)

	fmt.Printf("pi: %v, %T\n", pi, pi)
	fmt.Printf("ageLimit: %v, %T\n", ageLimit, ageLimit)
	fmt.Printf("defaultCountry: %v, %T\n", defaultCountry, defaultCountry)
	fmt.Printf("isDebuggingEnabled: %v, %T\n", isDebuggingEnabled, isDebuggingEnabled)
	fmt.Printf("_favChar: %v, %T\n", _favChar, _favChar)

	// variables
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	i := 7 // short declaration
	i = i + 9
	fmt.Println(i)
	i += 10
	fmt.Println(i)

	var t, u, s = 20, 30, "thi is a string"
	fmt.Println(t, u, s)
	x, y := 5, 12
	fmt.Println(x, y)

	g := 365.4
	d := g + 9
	fmt.Println(d)
	var e = int(g + 9)
	fmt.Printf("e: %v, %T\n", e, e)

	// fmt.Println(int(pi))
	h := pi
	fmt.Println(int(h))
	const pi2 = 3.0
	fmt.Println(int(pi2))
	const huge = 184467440737095516151844674407370955161518446744073709551615
	var z = huge / 1844674407370955161518446744073709551615184467440
	fmt.Println(z)
	fmt.Println(float64(huge))
}
func foo() {
	fmt.Println(pi)
}
