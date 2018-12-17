package main

import "fmt"

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
