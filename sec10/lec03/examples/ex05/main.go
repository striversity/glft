package main

import (
	"fmt"
	"log"
	"os"
)

// Section 10 - Lecture 03 : os.File

func main() {
	var fn = "data.txt"
	if len(os.Args) > 1 {
		fn = os.Args[1]
	}

	// opening and closing a file
	// ----
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// reading bytes from a file
	// ----
	b := make([]byte, 500)
	n, err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", string(b[:n]))
}
