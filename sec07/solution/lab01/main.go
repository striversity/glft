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
	result      map[string]int // word(string) -> count(int)
	wg          sync.WaitGroup
	resultMutex sync.Mutex
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

	wg.Wait() // wait until all files are processed
	elapse := time.Since(start)
	printResult()
	fmt.Printf("Total time: %v\n", elapse)
}

// processFile takes a filename 'fn' and reads each line of text. For
// each line, it call countWords() to count the number of words in the line.
func processFile(fn string) {
	wg.Add(1)
	go func() {
		log.Infof("Processing file %v", fn)
		fr, err := input.NewFileReader(fn)
		if err != nil {
			log.Warnf("Unable to open file %v: %v", fn, err)
			return
		}

		for fr.Scan() {
			l := fr.Text()
			// countWords splits a line of text into words. Each word is then added to result
			for _, pc := range replaceChars {
				l = strings.Replace(l, string(pc), "", -1)
			}
			l = strings.ToLower(l)
			words := strings.Split(l, " ")
			for _, w := range words {
				// NOTE: Can't update map concurrently, need to synchronize access
				resultMutex.Lock() // could have enclose the 'for' loop in the Mutex too
				result[w] = result[w] + 1
				resultMutex.Unlock()
			}
		}
		wg.Done()
	}()
}

func printResult() {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
