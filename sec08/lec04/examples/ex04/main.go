// Section 08 - Lecture 04 : Pointers and Functions
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

const (
	HEADING1 = "Heading-1"
	HEADING2 = "Heading-2"
	HEADING3 = "Heading-3"
	HEADING4 = "Heading-4"
	docDataSize = 100e6
)

type (
	Document struct {
		Name          string
		Author        string
		CreatedOn     time.Time
		ModifiedOn    time.Time
		Data          [docDataSize]byte
		NewLineIndex  map[int]int      // line# -> buffer offset
		PageIndex     map[int]int      // page# -> buffer offset
		HeadingsIndex map[string][]int // heading level -> list of buffer offsets
		usedBytes     int
	}
)

func main() {
	// passing around large amount of data for manipulation
	// ----
	var aDoc = loadDoc("dummy source")
	start := time.Now()
	printDoc(aDoc)
	fmt.Printf("Size of aDoc object: %v\n", unsafe.Sizeof(aDoc))
	fmt.Printf("Word Count: %v\n", countWords(aDoc))
	capitalizeHeading(aDoc, HEADING1)
	printDoc(aDoc)
	listHeadings(aDoc, HEADING1)
	fmt.Printf("Runtime: %v\n", time.Since(start))
}
func countWords(doc Document) (words int) {
	fmt.Printf("Size of object passed to countWords(): %v\n", unsafe.Sizeof(doc))
	// ...
	return rand.Int() % 200000
}
func capitalizeHeading(doc Document, level string) {
	fmt.Printf("Size of object passed to capitalizeHeading(): %v\n", unsafe.Sizeof(doc))
	doc.ModifiedOn = time.Now()
	s := string(doc.Data[:doc.usedBytes])
	s = strings.ToTitle(s)
	var buf bytes.Buffer
	fmt.Fprintln(&buf, s)
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	// ...
}
func listHeadings(doc Document, level string) {
	fmt.Printf("Size of object passed to listHeadings(): %v\n", unsafe.Sizeof(doc))
	// ...
}
func printDoc(doc Document) {
	fmt.Printf("Doc Name: %v\n", doc.Name)
	fmt.Printf("Doc Created On: %v\n", doc.CreatedOn)
	fmt.Printf("Doc Modified On: %v\n", doc.ModifiedOn)
	fmt.Printf("%-6v %v", 1, string(doc.Data[:doc.usedBytes]))
}
func loadDoc(src string) (doc Document) {
	doc.Name = src
	doc.Author = "Jane Smith"
	doc.CreatedOn = time.Now()
	doc.ModifiedOn = time.Now()
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "capter 1")
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	return
}
