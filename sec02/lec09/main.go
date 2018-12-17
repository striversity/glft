// Section 02 - Lecture 09 : Functions - Part 2
package main

import (
	"fmt"
)

func swap(s, t int) (int, int) {
	return t, s
}

func swap2(s, t int) (x int, y int) {
	x = t
	y = s
	return
}
func sum(v ...int) {
	fmt.Printf("v values are: %v\n", v)
	fmt.Printf("v's type: %T\n", v)
}
func formatter(s string, f ...float64) {
	///
}

var foo = func() {
	fmt.Println("I am the func foo()")
}
var hi = func(s string, c int) {
	for x := 0; x < c; x++ {
		fmt.Printf("Hi, %v\n", s)
	}
}
/* GIVEN - The functions:
db1 = [...]
db2 = [...]
getNumOfItemsDb1() int - which returns the nubmer of items in db1
getItemCostDb1(int) float64 - which returns the cost of an it in db1
NOTE: There are functions getNumOfItemsDb2() and getItemCostDb2() which operates on db2
*/

// stats - calculates some simple stats for a list of numbers
func stats(
	getNumOfItems func() int,
	getItemCost func(int) float64) {
	/*
		getNumOfItems() returns the number of items in the database
		getItemCost(itemID) returns the cost of an item,
		  where itemID is 1 to getNumOfItems(), 0.0 otherwise
	*/
	var itemCount = getNumOfItems()
	var totalCost, avgCost, maxCost, minCost float64
	for x := 1; x <= itemCount; x++ {
		c := getItemCost(x)
		totalCost += c
		if c > maxCost {
			maxCost = c
		}
		if c < minCost || minCost == 0.0 {
			minCost = c
		}
	}
	avgCost = totalCost / float64(itemCount)
	fmt.Printf("%v items processed\n", itemCount)
	fmt.Printf("Total price: $%.2f\n", totalCost)
	fmt.Printf("Average price: $%.2f\n", avgCost)
	fmt.Printf("Max price: $%.2f\n", maxCost)
	fmt.Printf("Min price: $%.2f\n", minCost)
}

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
	stats(getNumOfItemsDb1, getItemCostDb1)
	stats(getNumOfItemsDb2, getItemCostDb2)
	
}



// ------- DON"T LOOK BELOW HERE
var db1 = map[int]float64{
	1: 340.77, 2: 883.66, 6: 1411.14, 9: 1639.36, 10: 1481.56,
	3: 1084.18, 4: 577.13, 5: 796.16, 7: 224.36, 8: 1138.59,
}
var db2 = map[int]float64{
	1: 1098.12, 2: 519.92, 3: 628.44, 4: 126.35, 5: 697.27,
	6: 852.28, 7: 1576.11,
}

// getNumOfItemsDb1 returns the totalNumber of items in database 1
func getNumOfItemsDb1() int {
	return len(db1)
}

/* getItemCostDb1 returns the cost of an item identified by 'itemID',
where 'itemID' is 1 to getNumofItemsDb1() */
func getItemCostDb1(itemID int) (c float64) {
	c = db1[itemID]
	return
}

// getNumOfItemsDb2 returns the totalNumber of items in database 2
func getNumOfItemsDb2() int {
	return len(db2)
}

/* getItemCostDb2 returns the cost of an item identified by 'itemID',
where 'itemID' is 1 to getNumofItemsDb2() */
func getItemCostDb2(itemID int) (c float64) {
	c = db2[itemID]
	return
}
