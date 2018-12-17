// Section 03 - Excercise 05

package main

import (
	"fmt"
	"log"
)

func main() {
	testMyCopyStringSlice()
	var s []string
	s = myAppendStringToSlice(s, "I")
	s = myAppendStringToSlice(s, "love", "programming", "in", "Golang!")
	fmt.Println(s)  // [I love programming in Golang!]
}

// TODO 1 - implemenat myCopyStringSlice()
func myCopyStringSlice(dst []string, src []string) {
    // copy the minimum of len(dest) and len(src)
    num := min(len(dst), len(src))
    for i := 0; i < num; i++ {
        dst[i] = src[i]
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// TODO 2 - Impelment your custom []stirng append function myAppendStringToSlice()
func myAppendStringToSlice(dst []string, s ...string) []string {
	dLen := len(dst)
	sLen := len(s)
	var tmp = dst // assume dst is big enough
	if dLen < sLen {
		// destination not big enough, create a new slice that is big enough
		tmp = make([]string, dLen+sLen)
		// must copy old data in the slice
		copy(tmp, dst)
	}
	// do append operation
	copy(tmp[dLen:], s)
	return tmp
}

// --- NO need to look below here
// TODO 1 - Implement your custom []string copy function myCopyStringSlice()
func testMyCopyStringSlice() {
	s0 := []string{"To", "be", "or", "not", "to", "be.", "That", "is", "the", "question."}
	var s1 []string
	myCopyStringSlice(s1, s0)
	s2 := make([]string, 4)
	expected := make([]string, 4)
	myCopyStringSlice(s2, s0[0:0])
	copy(expected, s0[0:0])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
	myCopyStringSlice(s2, s0[4:])
	copy(expected, s0[4:])
	for i, v := range s2 {
		if v != expected[i] {
			log.Fatalf("test in testMyCopyStringSlice() failed, expected: %v, got: %v", expected, s2)
		}
	}
}
