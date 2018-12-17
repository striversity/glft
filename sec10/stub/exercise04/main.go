// Section 10 - Exercise 04 : Simple int/float to string conversion
package main

import (
	"io"
	"os"
)

var (
	digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func main() {
	io.WriteString(os.Stdout, intToStr(2010)+"\n")
	io.WriteString(os.Stdout, intToStr(1971/07/07)+"\n") // 40
	io.WriteString(os.Stdout, intToStr(29093423)+"\n")
	io.WriteString(os.Stdout, f64ToStr(9.15)+"\n")
	io.WriteString(os.Stdout, f64ToStr(11.04)+"\n")
}

// intToStr converts an int to a string without any help from any packages
func intToStr(i int) (s string) {
	// TODO 1
	return
}
// f64ToStr converts a float64 to a string without any help from any packages
func f64ToStr(f float64) (s string) {
	// TODO 2
	return
}
