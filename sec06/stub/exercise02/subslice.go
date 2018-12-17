package main

import (
	"math/rand"
	"time"
)

const (
	numTests = 10
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

// subSlice returns a slice of the 'data' such that 'data' can be considered broken up
// into subsets for 'n' subset. The parameter 'i' specifies which 0 based
// subset of the data to return. A nil slice is return if 'i' or 'n' is negative, or if
// 'data' is nil
func subSlice(data []int, n, i int) (out []int) {
	l := len(data)
	if l == 0 || 0 > i || 0 > n {
		return
	}
	maxSize := l / n
	if 0 != (l % n) {
		maxSize++ // increment the size by 1 if it doesn't divide evenly
	}
	offset := i * maxSize
	if offset > l {
		return
	}
	actSize := min(maxSize, l-offset)
	// log.Infof("setSlice(): len: %v, n: %v, i: %v, max size, %v, actual size: %v", l, n, i, maxSize, actSize)
	out = data[offset : offset+actSize]
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
