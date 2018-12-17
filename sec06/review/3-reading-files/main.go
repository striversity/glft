// Section 06 - Lab - Reading Text Files
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/striversity/glft/shared/input"
)

func main() {
	// Reading lines from a text file
	// 1. open file
	fr, err := input.NewFileReader(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to read file %v : %v", os.Args[1], err)
	}
	// 2. loop reading a line until end of file
	for fr.Scan() {
		s := fr.Text()
		//  2.1 process line
		fmt.Printf("Got line: %v\n", s)
	}
	// 3. close file
	fmt.Println("Done reading data")
}
