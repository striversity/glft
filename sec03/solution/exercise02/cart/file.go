/*
TODO 2 - Calculate discount and sales tax:
	a) Initializes an array representing the total purchase of 20 shoppers.
	   - TIP: Use sub-package 'cart.getTotal()'
	b) For each shopper, calculate a 5% discount and 8% sales tax
	c) Print the results in some meaningful way
*/
package cart

import (
	"fmt"

)

const numShoppers = 20

type dataSet [numShoppers]Currency

func Print() {
	cartTotals := getCartTotals()
	discountTotals := addDiscount(cartTotals, 5.0)
	totals := addSalesTax(discountTotals, 8.0)

	// https://golang.org/pkg/fmt/#hdr-Printing
	fmt.Printf("%-15v%-10v%-15v%-10v\n", "Shopper ID", "Inital", "5% Discounted", "Final")
	for i, c := range totals {
		fmt.Printf("%-15v%-10v%-15v%-10v\n", (i + 1), cartTotals[i], discountTotals[i], c)
	}
}
func addSalesTax(totals dataSet, tax float32) dataSet {

	for i := range totals {
		totals[i] += (totals[i] * (Currency(tax) / 100.0))
	}
	return totals
}
func addDiscount(totals dataSet, discount float32) dataSet {
	for i := range totals {
		totals[i] -= (totals[i] * (Currency(discount) / 100.0))
	}
	return totals
}
func getCartTotals() dataSet {
	var d dataSet
	for i := range d {
		d[i] = getTotal()
	}
	return d
}
