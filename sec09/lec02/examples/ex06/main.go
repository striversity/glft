// Section 09 - Lecture 02 : Implementing an interface
package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

const (
	HEADING1    = "Heading-1"
	HEADING2    = "Heading-2"
	HEADING3    = "Heading-3"
	HEADING4    = "Heading-4"
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
	var aDoc, _ = loadDoc("dummy source")
	start := time.Now()
	aDoc.printDoc()
	fmt.Printf("Size of aDoc object: %v\n", unsafe.Sizeof(aDoc))
	fmt.Printf("Word Count: %v\n", aDoc.countWords())
	aDoc.capitalizeHeading(HEADING1)
	aDoc.printDoc()
	aDoc.listHeadings(HEADING1)
	fmt.Printf("Runtime: %v\n", time.Since(start))
}
func (doc *Document) countWords() (words int) {
	if doc == nil {
		return
	}
	fmt.Printf("Size of object passed to countWords(): %v\n", unsafe.Sizeof(doc))
	// ...
	return rand.Int() % 200000
}
func (doc *Document) capitalizeHeading(level string) {
	if doc == nil {
		return
	}
	fmt.Printf("Size of object passed to capitalizeHeading(): %v\n", unsafe.Sizeof(doc))
	doc.ModifiedOn = time.Now()
	s := string(doc.Data[:doc.usedBytes])
	s = strings.ToTitle(s)
	var buf bytes.Buffer
	fmt.Fprintln(&buf, s)
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	// ...
}
func (doc *Document) listHeadings(level string) {
	fmt.Printf("Size of object passed to listHeadings(): %v\n", unsafe.Sizeof(doc))
	// ...
}
func (doc *Document) printDoc() {
	if doc == nil {
		return
	}
	fmt.Printf("Doc Name: %v\n", doc.Name)
	fmt.Printf("Doc Created On: %v\n", doc.CreatedOn)
	fmt.Printf("Doc Modified On: %v\n", doc.ModifiedOn)
	fmt.Printf("%-6v %v", 1, string(doc.Data[:doc.usedBytes]))
}
func loadDoc(src string) (doc *Document, err error) {
	doc = new(Document)
	if doc == nil { // can't alloate
		return nil, errors.New("Can't allocate new Document")
	}
	doc.Name = src
	doc.Author = "Jane Smith"
	doc.CreatedOn = time.Now()
	doc.ModifiedOn = time.Now()
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "capter 1")
	doc.usedBytes, _ = buf.Read(doc.Data[:100])
	return
}
