package main

import "fmt"

func main() {
	// using fmt.Sscanln - parsing int, float, and strings from a string
	// ----
	input := "Jane Doe\n35\n001-11-2233\n5.6"
	var i int
	var f float64
	var s0, s1, s2 string
	fmt.Sscanln(input, &s0, &s1, &i, &s2, &f)
	fmt.Printf("Got int i: %v\n", i)
	fmt.Printf("Got float64 f: %v\n", f)
	fmt.Printf("Got string s0: '%v'\n", s0)
	fmt.Printf("Got string s1: '%v'\n", s1)
	fmt.Printf("Got string s2: '%v'\n", s2)
}
