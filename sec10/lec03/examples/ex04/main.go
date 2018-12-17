package main

import (
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

}
