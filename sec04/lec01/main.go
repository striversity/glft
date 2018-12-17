// Section 04 - Lecture 01 : Maps
package main

import (
	"fmt"
)

func main() {
	// review declaring a slice
	var s []int
	if s != nil {
		fmt.Printf("s: %v, len: %v\n", s, len(s))
	}
	// declaring a map
	var m map[int]string
	fmt.Printf("m: %v, len: %v\n", m, len(m))
	// create a map using the make()
	// m[98] = "A+" -- illegal, can't store in 'nil' map
	m = make(map[int]string)
	fmt.Printf("m: %v, len: %v\n", m, len(m))
	// storing values in a map
	m[98] = "A+"
	m[95] = "A"
	m[-1] = "invalid"
	fmt.Printf("m: %v, len: %v\n", m, len(m))
	// retrieving values from a map
	fmt.Printf("m[98]: %v\n", m[98])
	// using non-existing key in a map
	fmt.Printf("key doesn't exisst: m[70]: %v\n", m[70])
	m1 := make(map[string][]string)
	m1["Smith"] = make([]string, 4)
	m1["Smith"][0] = "Mary"
	m1["Smith"][1] = "John"
	m1["Smith"][2] = "Pete"
	m1["Smith"][3] = "Anne"
	fmt.Printf("The Smith family: %v\n", m1["Smith"])
	jones := m1["Jones"]
	fmt.Printf("The Jones family: %v\n", jones)
	if jones == nil {
		fmt.Println("Confirmed, jones is nil")
	}
	// checking if a key existed during lookup
	if _, ok := m1["Jones"]; ok == false {
		fmt.Println("Jones is not in our map")
	}
	// length of a map
	fmt.Printf("number of elements in map: %v\n", len(m1))

	// map iteration
	for k, v := range m {
		fmt.Printf("m[%v] = %v\n", k, v)
	}

	// delete elements from a map
	delete(m, -1)
	fmt.Printf("m: %v\n", m)

	var a [10]int
	var a2 = a
	a[0] = -1
	a2[0] = 2
	fmt.Printf("a: %v, a2: %v\n", a, a2)
	var m2 = m
	fmt.Printf("m: %v, m2: %v\n", m, m2)
	m2[70] = "C+"
	fmt.Printf("m: %v, m2: %v\n", m, m2)
}
