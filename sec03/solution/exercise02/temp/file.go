/*
TODO 1 - Stats on temperature values
	1) declare an array of about a few hundred 'int'
	2) assign a random integer between -20 and 120 to data[i]
		- TIP: you can use input.GetRandInt(low, max) int
	3) print sorted list of tempature data
	4) print that max tempature
	5) print the min tempature
	6) print the average tempature

*/
package temp

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

const dataSize = 10

type dataSet [dataSize]int

var data dataSet

func init() {
	for i := 0; i < len(data); i++ {
		// TODO 1.a - assign a random integer betwen -20 and 120 to data[i]
		// TIP: you can use input.GetRandInt(low, max) int
		data[i] = input.GetRandInt(-20, 120)
	}
}

func Print() {
	// print some stats about the data
	fmt.Println("Data: ", data)
	sdata := sort(data)
	fmt.Println("Sorted: ", sdata)
	fmt.Println("Max: ", max(sdata))
	fmt.Println("Min: ", min(sdata))
	fmt.Println("Avg: ", avg(data))
}

/**
TODO 1.b - sort an array of integers
algorithm:
  1. for each element in the array at 'i'
  2. check if there is another element in the rest of the array that is
     less than or equal to the element at 'i'
	  - iterator j will be used to check the rest of the array from 'i+1'
	    to the last element of the array
*/
func sort(d dataSet) dataSet {
	for i := 0; i < len(d); i++ {
		for j := i + 1; j < len(d); j++ {
			if d[j] <= d[i] {
				// TODO 2 - Swap the value of d[i] with d[j]
				d[i], d[j] = d[j], d[i]
			}
		}
	}

	return d
}

// max returns the maximum value of a sorted data set
// NOTE: data set MUST be sorted
func max(d dataSet) int {
	return d[len(d)-1]
}

// min returns the minimun values of a sorted data set
// NOTE: data set MUST be sorted
func min(d dataSet) int {
	return d[0]
}

func sum(d dataSet) int {
	s := d[0]
	for i := 1; i < len(d); i++ {
		s += d[i]
	}

	return s
}

func avg(d dataSet) int {
	return (sum(d) / len(d))
}
