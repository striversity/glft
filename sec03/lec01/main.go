// Section 03 - Lecture 01 : // Declaring and Using Arrays
package main

import (
	"fmt"
)

type (
	Currency float64
)

func (c Currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
func main() {
	// world before arrays
	var price0 Currency = 9.23
	var price1 Currency = 0.59
	var price2 Currency = 12.84
	var price3 Currency = 5.95
	var price4 Currency = 16.37
	var price5 Currency = 1.35

	fmt.Println(price0, price1, price2, price3, price4, price5)
	totalPrice := (price0 + price1 + price2 +
		price3 + price4 + price5)
	avgPrice := totalPrice / 6
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// declaring an array
	var prices [6]Currency
	fmt.Println(prices)
	prices[0] = 9.23
	prices[1] = 0.59
	prices[2] = 12.84
	prices[3] = 5.95
	prices[4] = 16.37
	prices[5] = 1.35
	//prices[6] = 1.35 - illegal, causes error
	fmt.Println(prices)

	totalPrice = (prices[0] + prices[1] + prices[2] +
		prices[3] + prices[4] + prices[5])
	avgPrice = totalPrice / 6
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// using 'for' loop to iterate over an Array
	totalPrice = 0
	for i := 0; i < 6; i++ {
		totalPrice = totalPrice + prices[i]
	}
	avgPrice = totalPrice / 6
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// using the len() function - builtin
	totalPrice = 0
	for i := 0; i < len(prices); i++ {
		totalPrice = totalPrice + prices[i]
	}
	avgPrice = totalPrice / Currency(len(prices))
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// iterate over an array using 'range' operator - builtin
	totalPrice = 0
	for i := range prices {
		totalPrice += prices[i]
	}
	avgPrice = totalPrice / Currency(len(prices))
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// iterate over an array using 'range' operator , index and value
	totalPrice = 0
	for i, v := range prices {
		fmt.Printf("prices[%v] = %v\n", i, v)
		totalPrice += v
	}
	avgPrice = totalPrice / Currency(len(prices))
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// iterate over an array using 'range' operator , index and value
	totalPrice = 0
	for _, v := range prices {
		totalPrice += v
	}
	avgPrice = totalPrice / Currency(len(prices))
	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)

	// var x = 5
	var prices2 = [60]Currency{
		9.23,
		0.59,
		12.8,
		5.95,
		16.3,
		1.35,
	}

	fmt.Println(prices2)
}
