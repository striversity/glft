package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	
	// fmt.Fscanln - reading from a file
	// ----
	var i int
	var f float64
	var s0, s1 string
	fmt.Fscanln(input, &s0, &i, &s1, &f)
	fmt.Printf("Got int i: %v\n", i)
	fmt.Printf("Got float64 f: %v\n", f)
	fmt.Printf("Got string s0: '%v'\n", s0)
	fmt.Printf("Got string s1: '%v'\n", s1)
}
