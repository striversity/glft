// Section 09 - Lecture 03 : Method Sets
package main

import "fmt"

type (
	Person struct {
		fname string
		lname string
		age   uint8
	}
)

func main() {
	// method sets for T and *T using variable of type *T
	// ----
	p := &Person{"Mark", "Smith", 35}
	fmt.Printf("Mark's name: %v\n", p.Name()) // -> *(p).Name()
	fmt.Printf("Mark's age: %v\n", p.Age())   // -> *(p).Age()
	p.SetAge(p.Age() + 1)
	fmt.Printf("Mark's age: %v\n", p.Age()) // -> *(p).Age()
}

func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}
func (p Person) Age() uint8 {
	return p.age
}
func (p *Person) SetAge(a uint8) {
	if a <= 150 && a > p.age {
		fmt.Printf("Chaning age from %v to %v\n", p.age, a)
		p.age = a
	}
}
