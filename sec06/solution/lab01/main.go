// Section 06 - Lab 01 : Work Count (iterative)
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/shared/input"
)

const (
	replaceChars = "`" + `~!@#$%^&*()-_+=[{]}\|;:'",<.>/?`
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
	fr, err := input.NewFileReader(fn)
	if err != nil {
		log.Warnf("Unable to open file %v: %v", fn, err)
		return
	}

	for fr.Scan() {
		countWords(fr.Text())
	}
}

// countWords splits a line of text into words. Each word is then added to result
func countWords(l string) {
	for _, pc := range replaceChars {
		l = strings.Replace(l, string(pc), "", -1)
	}
	l = strings.ToLower(l)
	words := strings.Split(l, " ")
	for _, w := range words {
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
