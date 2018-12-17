// Section 02 - Lecture 11 : Packages - Part 1
package main

import (
	"fmt"
	 "github.com/verrol/glft/sec02/lec11/db"
)

/* GIVEN - The functions:
db1 = [...]
db2 = [...]
getNumOfItemsDb1() int - which returns the nubmer of items in db1
getItemCostDb1(int) float64 - which returns the cost of an it in db1
NOTE: There are functions getNumOfItemsDb2() and getItemCostDb2() which operates on db2
*/


func main() {
	a, b := 9, 15
	fmt.Printf("a, b: %v, %v\n", a, b) // 9, 15
	a, b = swap(a, b)
	fmt.Printf("a, b: %v, %v\n", a, b) // 15, 9

	a, b = 11, 4
	fmt.Printf("a, b: %v, %v\n", a, b) // 11, 4
	a, b = swap2(a, b)
	fmt.Printf("a, b: %v, %v\n", a, b) // 4, 11

	// varadic function
	sum()            // 0 parameters
	sum(1)           // 1 parameter
	sum(3, 5, 9, 10) // multiple parameters
	formatter("format 1")
	formatter("format 2", 1.0)
	formatter("format 3", 3.13, 9.0, -9.0)
	formatter("format 4", 1.0, 93.9, 89.1, 173.0)
	formatter("format 5", -13, 53, 544.0, 342, 1343, 0.3)

	// ananymous functions
	foo()
	hi("World!", 2)
	hi("Verrol", 4)
	fmt.Printf("foo's type: %T\n", foo)
	fmt.Printf("hi's type: %T\n", hi)
	var hi2 func(string, int)
	fmt.Printf("hi2's value: %v\n", hi2)
	hi2 = hi
	fmt.Printf("hi2's value: %v\n", hi2)
	hi2("Earth!", 5)

	// call stats
	stats(db.GetNumOfItemsDb1, db.GetItemCostDb1)
	stats(db.GetNumOfItemsDb2, db.GetItemCostDb2)
	
}