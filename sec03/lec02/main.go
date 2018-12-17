// Section 03 - Lecture 02 : // Arrays and Functions
package main

import (
	"fmt"

	"github.com/verrol/glft/shared/input"
)

const dataSize = 10

type dataSet [dataSize]int

func main() {
	var data dataSet
	for i := 0; i < len(data); i++ {
		data[i] = input.GetRandInt(-10, 20)
	}
	fmt.Printf("Data set: %v\n", data)
	fmt.Printf("Minimum value: %v\n", min(data))
	sdata := sort(data)
	fmt.Printf("Sorted data set: %v\n", sdata)
}
func sort(d dataSet) dataSet {
	for i := range d {
		for j := i + 1; j < len(d); j++ {
			if d[j] < d[i] {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
	return d
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
