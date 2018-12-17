// Section 06 - Lab 01 : Work Count (iterative)
package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	result map[string]int // word(string) -> count(int)
)

func main() {
	if 1 == len(os.Args) {
		log.Fatal("No files to process")
	}

	inputFiles := os.Args[1:]
	result = make(map[string]int)

	start := time.Now()
	log.Infof("Processing %v files for input", len(inputFiles))
	for _, fn := range inputFiles {
		processFile(fn)
	}

	elapse := time.Since(start)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

// processFile takes a filename 'fn' and reads each line of text. For
// each line, it call countWords() to count the number of words in the line.
func processFile(fn string) {
	log.Infof("Processing file %v", fn)

}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
