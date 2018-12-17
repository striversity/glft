package cart

import (
	"math/rand"
	"time"
)

var r = rand.NewSource(time.Now().UnixNano())

// getTotal returns the total price of a shopper's cart
func getTotal() Currency {
	v := r.Int63()
	v = 1972 + (v % 45)
	c := Currency(v) / 11.4
	return c
}
