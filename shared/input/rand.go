package input

import (
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

// GetRandInt returns a number n, such that low <= n <= max
func GetRandInt(low, max int) int {
	if low >= max {
		return 0
	}
	span := max - low

	n := r.Int()
	return (low + (n % span))
}

// GetRandFloat returns a float64 number n between low <= n <= max
func GetRandFloat(low, max int64) float64 {
	if low >= max {
		return 0
	}
	span := max - low

	n := r.Float64() * float64(span)
	return (float64(low) + n)
}
