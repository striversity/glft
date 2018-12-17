/*
# Section 04 - Lab 02 : Word Count 2

## TODO 1 - Ingnore _noise words_ from your result

### Requirements

1. Use your solution to *Lab01* as starting point for *Lab02*
2. Create a list of words that shoudlnt' be included in the result.
    - Examples of _noise words_ are: 'on', 'a', 'the', 'are', 'in', 'of', etc.
	- TIP: Consider using *strings.ToLower()* or *strings.ToUpper()* to 
    filter out all forms of the noise words. For example, "on" vs "On" should be consisdered the same.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/striversity/glft/shared/input"
)

// TODO 1 - add additional 'noise' words to the list
var noiseWords = []string{
	"is", "to", "be", "in", "of", "for", "or", "not", "an", "are", "by",
}

var result map[string]int

func main() {
	fr, err := input.NewFileReader(os.Args[1])
	if nil != err {
		log.Fatal(err)
	}

	result = make(map[string]int)

	var line string
	var words []string

	for nil == err {
		line, err = fr.ReadLine()
		line = strings.TrimSpace(line)
		words = strings.Split(line, " ")
		processWords(words)
	}

	printResult()
}

func processWords(words []string) {
	skip := false
	for _, w := range words {
		for _, t := range noiseWords {
			// compare case insensitive word
			if strings.ToLower(w) == strings.ToLower(t) {
				skip = true
				break
			}
		}

		if !skip {
			result[w] = result[w] + 1
		}
		skip = false
	}
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
