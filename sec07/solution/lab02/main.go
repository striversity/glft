// Section 06 - Lab 02 : Work Count (concurrent)
package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/shared/input"
)

const (
	replaceChars = "`" + `~!@#$%^&*()-_+=[{]}\|;:'",<.>/?`
)

var (
	results chan map[string]int // word(string) -> count(int)
	wg      sync.WaitGroup
)

func main() {
	if 1 == len(os.Args) {
		log.Fatal("No files to process")
	}

	inputFiles := os.Args[1:]
	results = make(chan map[string]int, len(inputFiles))

	start := time.Now()
	log.Infof("Processing %v files for input", len(inputFiles))
	for _, fn := range inputFiles {
		out := processFile(fn) // generator
		countWords(out)        // sink
	}

	wg.Wait() // wait until all files are processed
	elapse := time.Since(start)
	close(results)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

// processFile is a generator which takes a filename 'fn' and reads each line of text. For
// each line, it sends it on the channel 'out'.
func processFile(fn string) (out <-chan string) {
	wg.Add(1)
	ch := make(chan string)
	go func() {
		defer wg.Done()
		out = ch
		log.Infof("Processing file %v", fn)
		fr, err := input.NewFileReader(fn)
		if err != nil {
			log.Warnf("Unable to open file %v: %v", fn, err)
			return
		}

		for fr.Scan() {
			ch <- fr.Text()
		}
		close(ch)
	}()
	return ch
}

// countWords splits a line of text into words. Each word is then added to result
func countWords(in <-chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := make(map[string]int)
		for l := range in {
			for _, pc := range replaceChars {
				l = strings.Replace(l, string(pc), "", -1)
			}
			l = strings.ToLower(l)
			words := strings.Split(l, " ")
			for _, w := range words {
				result[w] = result[w] + 1
			}
		}
		results <- result
	}()
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for result := range results {
		for w, c := range result {
			fmt.Printf("%-10v%s\n", c, w)
		}
	}
}
