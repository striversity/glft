// Section 02 - Lecture 07 : The "for" statement
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalMessages = 5

func main() {
	// 'for' statement
	x := 1
	for x <= totalMessages {
		fmt.Printf("This is message %v of %v\n", x, totalMessages)
		x++ // // x += 1 ==> x = x + 1
	}

	// 'for' statement with short var declaration
	for x := 1; x <= totalMessages; x++ {
		fmt.Printf("This is message %v of %v\n", x, totalMessages)
	}

	// calculation with sample data
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

	// count how many items cost $1000.00 or more
	var numMore1000 int
	const cutOffPrice = 1000.0
	for x := 1; x <= itemCount; x++ {
		c := getItemCost(x)
		if c < cutOffPrice {
			continue
		}
		numMore1000++
	}
	fmt.Printf("There are %v items which are $%.2f or more.\n", numMore1000, cutOffPrice)
}

// ------- DON"T LOOK BELOW HERE
const dataPoints = 10

var db map[int]float64

// getNumOfItems returns the totalNumber of items in our database
func getNumOfItems() int {
	return dataPoints
}

// getItemCost returns the cost of an item identified by 'itemID', where 'itemID' is 1 to getNumofItems()
func getItemCost(itemID int) (c float64) {
	c = db[itemID]
	return
}

// init initializes our database of items and prices
func init() {
	db = make(map[int]float64)
	for i := 1; i <= dataPoints; i++ {
		rand.Seed(time.Now().Unix() + rand.Int63())
		// get a random float value
		v := rand.Uint32()
		v %= 200000                // limit range to no more than 200,000
		db[i] = float64(v) / 100.0 // results in 2 decimal places, price between $0 to $2,000.00
	}
}
