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
	// using interface variable to enforce function contract
	// ----
	var e Printer
	Println(e,
		ID(19721104),
		SSN("019-72-1104"),
		&Person{name: "Jane Doe", age: 35, ssn: SSN("019-72-1104")})
}
func (p *Person) Print() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[Person] Name: %v, Age: %v, SSN: %v", p.name, p.age, p.ssn.Print())
}
func (ssn SSN) Print() string {
	return fmt.Sprintf("[SSN] %v", string(ssn))
}
func (id ID) Print() string {
	return fmt.Sprintf("[ID] %v", uint64(id))
}

// Println calls the 'Print() string' method of the value e
func Println(e ...Printer) {
	fmt.Println("[main.Println]")
	for i := range e {
		if e[i] == nil {
			fmt.Printf(" parameter[%v] is <nil>\n", i)
			continue
		}
		fmt.Printf(" parameter[%v]'s .Print() value: %v\n", i, e[i].Print())
	}
}
