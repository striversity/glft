// Section 03 - Lecture 05 : CECS (creating, expanding,
// copying, and shrinking) Slices
package main

import (
	"fmt"
)

type (
	Curreny float64
)

func main() {
	var s0 []int // slice of 'int' variable
	fmt.Printf("s0: %v, len: %v, type: %T\n", s0, len(s0), s0)
	var ar [10]int
	s0 = ar[:] // create slice from an array
	fmt.Printf("s0: %v, len: %v, type: %T\n", s0, len(s0), s0)

	// creating slices at runtime using make()
	s1 := make([]int, 10)
	s1[0] = -1
	fmt.Printf("s1: %v, len: %v, type: %T\n", s1, len(s1), s1)
	s2 := make([]int, 5)
	s2[0] = 100
	fmt.Printf("s2: %v, len: %v, type: %T\n", s2, len(s2), s2)
	s3 := make([]Curreny, 7)
	s3[0], s3[2], s3[5] = 3.19, 6.09, 19.86
	fmt.Printf("s3: %v, len: %v, type: %T\n", s3, len(s3), s3)

	// expaning slices
	var s4 []string
	fmt.Printf("s4: %v, len: %v\n", s4, len(s4))
	s4 = append(s4, "New York")
	fmt.Printf("s4: %v, len: %v\n", s4, len(s4))
	s4 = append(s4, "Chicago")
	fmt.Printf("s4: %v, len: %v\n", s4, len(s4))
	s4 = append(s4, "Charlotte", "San Francisco", "Berlin")
	fmt.Printf("s4: %v, len: %v\n", s4, len(s4))

	// copy slices, NOTE: *shallow* copy
	var s5 = s4
	s5[0] = "Georgetown"
	fmt.Printf("s4: %v, len: %v\n", s4, len(s4))
	fmt.Printf("s5: %v, len: %v\n", s5, len(s5))
	
	// copy slices - DEEP copy
	s6:= make([]string, 3)
	copy(s6, s5)
	fmt.Printf("s6: %v, len: %v\n", s6, len(s6))
	s7:= make([]string, 7)
	s7[4],s7[5],s7[6] = "Hello", "World", "Nice!"
	copy(s7, s5)
	fmt.Printf("s7: %v, len: %v\n", s7, len(s7))

	var s8 []string
	s8 = append(s8, s7...)
	fmt.Printf("s8: %v, len: %v\n", s8, len(s8))
	
	// shriking slices
	s9 := s8[2:4]
	fmt.Printf("s9: %v, len: %v\n", s9, len(s9))
	
	var s10 []string
	s10 = append(s10, s8[4:]...)
	fmt.Printf("s10: %v, len: %v\n", s10, len(s10))
}
