// Section 02 - Lecture 10 : Types
package main

import (
	"fmt"
)

type (
	currency float64
	greeter  func() string
)

func formatCurrency(c float64) string {
	s := fmt.Sprintf("$%.2f", c)
	return s
}

func (c currency) String() string {
	s := fmt.Sprintf("$%.2f", float64(c))
	return s
}
func hi() string {
	return "Hi"
}
func goodMorning() string {
	return "Good Morning"
}
func say(greet greeter, s string) {
	fmt.Printf("%v, %v\n", greet(), s)
}
func main() {
	var price1 currency = 10.38
	var price2 currency = 29.69
	total := price1 + price2
	fmt.Printf("Total price: %v\n", total)
	a := avg(price1, price2)
	fmt.Printf("Avg price: %v, type of a: %T\n", a, a)
	formatCurrency(10.98)
	// formatCurrency(price1) - not allowed
	formatCurrency(float64(price1))

	say(hi, "World!")
	say(goodMorning, "World!")
}
func avg(v1 currency, v2 currency) (a currency) {
	a = (v1 + v2) / 2.0
	return
}
