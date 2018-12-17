// Section 09 - Lecture 01 : What is an interface?
package main

import "fmt"

type (
	Reader interface {
		Read()
	}
	Writer interface {
		Write() string
	}
	Circle interface {
		Radius() float64
		Area() float64
	}
	Rect interface {
		Width() float64
		Length() float64
		Area() float64
	}
)

func main() {
	// more examples of interface declaration
	// ----
	var dataSource Reader
	var printer Writer

	fmt.Printf("dataSource's value: %v, type: %T\n", dataSource, dataSource)
	fmt.Printf("printer's value: %v, type: %T\n", printer, printer)

}
