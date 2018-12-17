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
	// method sets for T and *T
	// ----
	mark := Person{"Mark", "Smith", 35}
	fmt.Printf("Mark's name: %v\n", mark.Name())
	fmt.Printf("Mark's age: %v\n", mark.Age())

	// using variable of type T to access *T method set
	mark.SetAge(mark.Age() + 1)                // -> m:= &mark; m.SetAge(mark.Age() + 1)
	fmt.Printf("Mark's age: %v\n", mark.Age())

}

func (p Person) Name() string {
	return fmt.Sprintf("%v, %v", p.lname, p.fname)
}
func (p Person) Age() uint8 {
	return p.age
}
func (p *Person) SetAge(a uint8) {
	if a <= 150 && a > p.age {
		fmt.Printf("Chaning age of %v from %v to %v\n", p.Name(), p.age, a)
		p.age = a
	}
}
