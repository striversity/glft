// Section 09 - Lecture 04 : interfaces and functions
package main

import "fmt"

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Printer interface {
		Print() string
	}
)

func main() {
	// interface as function return type
	// ----
	p := getValue()
	Println(p)
	p = getValue()
	Println(p)
}

func getValue() Printer {
	ch := make(chan int, 1)

	// send either 0 or 1 with equal probability
	select {
	case ch <- 1:
	case ch <- 0:
	}
	choice := <-ch

	if choice == 0 {
		return SSN("019-72-1104")
	}
	return ID(20100915)
}

func (id ID) Print() string {
	return fmt.Sprintf("[ID] %v", uint64(id))
}
func (ssn SSN) Print() string {
	return fmt.Sprintf("[SSN] %v", string(ssn))
}

// Println is package's main custom print func
func Println(e ...Printer) {
	fmt.Println("[main.Println]")
	for i := range e {
		fmt.Printf(" parameter[%v]'s .Print() value: %v\n", i, e[i].Print())
	}
}
