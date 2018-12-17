// Section 03 - Lecture 04 : // Slices and Functions
package main

import (
	"fmt"

	"github.com/verrol/glft/shared/input"
)

const dataSize = 10

type dataSet []int

func main() {
	var data dataSet
	var ar [dataSize]int
	data = ar[:]
	for i := 0; i < len(data); i++ {
		data[i] = input.GetRandInt(-10, 20)
	}
	fmt.Printf("Data set: %v\n", data)
	fmt.Printf("Minimum value: %v\n", min(data))
	sort(data)
	fmt.Printf("Sorted data set: %v\n", data)

	// smaller data set
	data = ar[5:]
	for i := 0; i < len(data); i++ {
		data[i] = input.GetRandInt(-10, 20)
	}
	fmt.Printf("Data set: %v\n", data)
	fmt.Printf("Minimum value: %v\n", min(data))
	sort(data)
	fmt.Printf("Sorted data set: %v\n", data)

	// very large dataset
	var vla [dataSize * 10000]int
	data = vla[:]
	for i := 0; i < len(data); i++ {
		data[i] = input.GetRandInt(-15000, 200000)
	}
	//fmt.Printf("Data set: %v\n", data)
	fmt.Printf("Minimum value: %v\n", min(data))
	//sort(data)
	fmt.Printf("Sorted data set: %v\n", data[100:120])

	// calling sum
	fmt.Printf("sum(): %v\n", sum())
	fmt.Printf("sum(): %v\n", sum(5))
	fmt.Printf("sum(): %v\n", sum(51, 21, -19, -142, 98))
	fmt.Printf("sum(): %v\n", sum(data[:10]...))
	ar2 := [...]int{-1, -10, 91, 86, -20}
	fmt.Printf("ar2 type: %T\n", ar2)
	fmt.Printf("sum(): %v\n", sum(ar2[:]...))
}
func sum(v ...int) (s int) {
	for _, v := range v {
		s += v
	}
	return
}
func sort(d dataSet) {
	for i := range d {
		for j := i + 1; j < len(d); j++ {
			if d[j] < d[i] {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
}
func min(d dataSet) int {
	m := d[0]
	for i := 1; i < len(d); i++ {
		if d[i] < m {
			m = d[i]
		}
	}
	return m
}
