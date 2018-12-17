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

func intToStr(i int) (s string) {
	for {
		r := i % 10
		i /= 10
		s = digits[r] + s
		if i == 0 { // finish when i = 0
			return
		}
	}
}
func f64ToStr(f float64) (s string) {
	w := int(f)
	s = intToStr(w)
	s += "."
	dec := f - float64(w)
	// fmt.Printf("%v after removing whole number: %v\n", f, dec)
	for {
		t:= int(dec * 10)
		s += intToStr(t)
		dec = (dec * 10) - float64(t)
		if dec < 0.001 {
			return
		}
	}
}
