// Section 08 - Lecture 04 : Pointers and Functions
package main

import (
	"fmt"
	"math/rand"
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
	}
)

func main() {
	// passing around large amount of data for manipulation
	// ----
	var aDoc Document
	fmt.Printf("Size of aDoc object: %v\n", unsafe.Sizeof(aDoc))
	fmt.Printf("Word Count: %v\n", countWords(aDoc))
	capitalizeHeading(aDoc, HEADING1)
	listHeadings(aDoc, HEADING1)
}
func countWords(doc Document) (words int) {
	fmt.Printf("Size of object passed to countWords(): %v\n", unsafe.Sizeof(doc))
	// ...
	return rand.Int() % 200000
}
func capitalizeHeading(doc Document, level string) {
	fmt.Printf("Size of object passed to capitalizeHeading(): %v\n", unsafe.Sizeof(doc))
	doc.ModifiedOn = time.Now()
	// ...
}
func listHeadings(doc Document, level string) {
	fmt.Printf("Size of object passed to listHeadings(): %v\n", unsafe.Sizeof(doc))
	// ...
}
