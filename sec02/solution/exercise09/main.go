package main

import (
	"fmt"
)

/* GIVEN - The functions:
db1 = [...]
db2 = [...]
getNumOfItemsDb1() int - which returns the nubmer of items in db1
getItemCostDb1(int) float64 - which returns the cost of an it in db1
NOTE: There are functions getNumOfItemsDb2() and getItemCostDb2() which operates on db2

/* TODO 1 - modify 'countAbove()' function to accept two additional 
parameters to represent the getNumberOfItems() and getItemCost() function */
var countAbove = func(cutOffPrice float64,
	getNumOfItems func() int,
	getItemCost func(int) float64) (count uint) {
	for i := 1; i <= getNumOfItems(); i++ {
		c := getItemCost(i)
		if c < cutOffPrice {
			continue
		}
		count++
	}
	return
}

func main() {
	// TODO 1 - count how many items cost $1000.00 or more in db1
	cutOff := 1000.0
	c := countAbove(cutOff, getNumOfItemsDb1, getItemCostDb1)
	fmt.Printf("There are %v items which cost $%.2f or more\n", c, cutOff)

	// TODO 2 - count how many items cost $ 500.00 or more in db2
	cutOff = 500
	c = countAbove(cutOff, getNumOfItemsDb2, getItemCostDb2)
	fmt.Printf("There are %v items which cost $%.2f or more\n", c, cutOff)
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
