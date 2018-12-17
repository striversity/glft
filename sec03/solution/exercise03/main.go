// Section 03 Exercise 03 - Declaring and Using Slices

package main

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

type (
	currency float64
)

func (c currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}

const (
	dataSize = 20
)

func main() {
	// TODO 1
	var db [dataSize]currency
	s := db[:]

	// -- NOTE: 'db' not used beyond this point

	// TODO 2 - Initialize the array with random currency values using 'input.GetRandFloat()'.
	for i := 0; i < dataSize; i++ {
		s[i] = currency(input.GetRandFloat(20, 100)) // prices between $20.00 and $100.00
	}

	// TODO 3 - Calculate total, max, min, avg
	var totalPrice, minPrice, maxPrice, avgPrice currency
	for _, v := range s {
		totalPrice += v
		if minPrice > v || minPrice == 0.0 {
			minPrice = v
		}
		if maxPrice < v {
			maxPrice = v
		}
	}
	avgPrice = totalPrice / dataSize

	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Min price: %v\n", minPrice)
	fmt.Printf("Max price: %v\n", maxPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)
}
