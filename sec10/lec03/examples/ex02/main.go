package main

import (
	"io"
	"log"
	"os"
)

// Section 10 - Lecture 03 : os.File

func main() {
	// create a file data.txt
	// ----
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("Hello, world!\n")
	io.WriteString(f, "This is yet another string\n")
}
