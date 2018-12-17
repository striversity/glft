package types

import "fmt"

type (
	Person struct {
		fname string
		lname string
		age   Age
	}
	Age uint8
)

func NewPerson(firstName, lastName string, age Age) (p Person) {
	p.age = age
	p.lname = lastName
	p.fname = firstName
	return
}
func (p Person) SetFirstName(s string) Person {
	if 0 < len(s) {
		p.fname = s
	}
	return p
}
func (p Person) Age(age Age) Age {
	return p.age
}
func (p Person) SetAge(age Age) Person {
	if age <= 150 {
		p.age = age
	}
	return p
}
func (p Person) String() string {
	s := fmt.Sprintf("%v, %v (%v)", p.lname, p.fname, p.age)
	return s
}
func (a Age) String() string {
	s := fmt.Sprintf("%v yrs", uint8(a))
	return s
}
