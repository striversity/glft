/*
# Section 04 - Lab 01 : Word Count 1

## TODO 1 - Count the occurrence of each _word_ in a text file

### Requirements

1. The filename is passed as an argument to the program
2. Use *input.NewFileReader()* to create an *input.FileReader* object.
    - The *input.FileReader* object has the method *ReadLine()*.
3. You will also need *strings.Split()* to break lines into words.
    - TIP: I also used *strings.TrimSpace()* in my solution.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/striversity/glft/shared/input"
)

var result map[string]int

func main() {
	fr, err := input.NewFileReader(os.Args[1])
	if nil != err {
		log.Fatal(err)
	}

	// TODO 1 - initialize result map
	result = make(map[string]int)

	var line string
	var words []string

	for nil == err {
		line, err = fr.ReadLine()
		// TODO 2 - remove newline charater from each line
		// TIP: checkout strings.TrimSpace()
		line = strings.TrimSpace(line)
		words = strings.Split(line, " ")
		processWords(words)
	}

	printResult()
}

func processWords(words []string) {
	for _, w := range words {
		// TODO 3 - update the map for each word
		result[w] = result[w] + 1
	}
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
