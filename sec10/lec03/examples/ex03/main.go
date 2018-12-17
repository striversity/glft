package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/striversity/glft/sec10/solution/exercise02/rw"
)

// Section 10 - Lecture 03 : os.File

func main() {
	// writing records to file data.txt
	// ----
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, _ := rw.NewRecordWriter(f)
	io.WriteString(w, "Hello, world!\n")
	io.WriteString(w, "This is yet another string\n")

	fmt.Printf("Write data file: %v\n", f.Name())
}
