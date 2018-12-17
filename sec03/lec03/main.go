// Section 03 - Lecture 03 : Declaring and Using Slices
package main

import (
	"fmt"
)

func main() {
	// var v type
	var a [10]int // array of 10 'int' values
	fmt.Printf("a: %v, type of a: %T\n", a, a)

	type currency float64
	var cart [10]currency // array of 10 'currecy' values
	fmt.Printf("cart: %v, type of cart: %T\n", cart, cart)

	// declaring a slice
	var s0 []int
	fmt.Printf("s0: %v\n", s0)
	fmt.Printf("s0 len: %v, type of s0: %T\n", len(s0), s0)

	nums := [10]int{12, 53, 86, 94, 75, 10, 2, 15, 37, 40}
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("nums len: %v, type of nums: %T\n", len(nums), nums)

	s1 := nums[:] //nums[0:len(nums)]
	fmt.Printf("s1 = nums[:]: %v\n", s1)
	fmt.Printf("s1 len: %v, type of s1: %T\n", len(s1), s1)

	s2 := nums[:5] // nums[0:5]
	fmt.Printf("s2 = nums[0:5]: %v\n", s2)
	fmt.Printf("s2 len: %v, type of s2: %T\n", len(s2), s2)

	s3 := nums[5:] // nums[5:len(nums)]
	fmt.Printf("s3 = nums[5:]: %v\n", s3)
	fmt.Printf("s3 len: %v, type of s3: %T\n", len(s3), s3)

	s4 := nums[3:7]
	fmt.Printf("s4 = nums[5:]: %v\n", s4)
	fmt.Printf("s4 len: %v, type of s4: %T\n", len(s4), s4)

	// modifying the underlying array via it's ***slices***
	for i := range s1 {
		fmt.Printf("s1[%v] = %v\n", i, s1[i])
		s1[i] = i
	}
	fmt.Printf("nums: %v\n", nums)

	for i := range s2 {
		fmt.Printf("s2[%v] = %v\n", i, s2[i])
		s2[i] = -1
	}
	fmt.Printf("nums: %v\n", nums)

	for i := range s3 {
		fmt.Printf("s3[%v] = %v\n", i, s3[i])
		s3[i] = 1
	}
	fmt.Printf("nums: %v\n", nums)

	for i := range s4 {
		fmt.Printf("s4[%v] = %v\n", i, s4[i])
		s4[i] = 0
	}
	fmt.Printf("nums: %v\n", nums)
}
