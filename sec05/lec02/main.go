// Section 05 - Lecture 02 : Initializtion, Nestings, and Anonymous Structs
package main

import (
	"fmt"
)

type (
	Address struct {
		street1 string
		street2 string
	}
	Age    uint8
	Person struct {
		firstName string
		lastName  string
		ssn       string
		age       Age
		address   Address
	}
	City struct {
		string
		State
	}
	State struct {
		name string
	}
	Country struct {
		name string
	}
	City2 struct {
		string
		State
		Country
	}
)

func main() {
	// initializing struct fields
	var a0 Address
	a0.street1 = "10 6th Ave."
	a0.street2 = "2nd Flr."
	fmt.Printf("%#v\n", a0)
	var a1 = Address{street1: "10 6th Ave.", street2: "2nd Flr."}
	fmt.Printf("%#v\n", a1)
	var a2 = Address{street2: "2nd Flr."}
	fmt.Printf("%#v\n", a2)
	var a3 = Address{"10 6th Ave.", "2nd Flr."}
	fmt.Printf("%#v\n", a3)
	var a4 = Address{street2: "2nd Flr.", street1: "10 6th Ave."}
	fmt.Printf("%#v\n", a4)
	var a5 = Address{
		street2: "2nd Flr.",
		street1: "10 6th Ave.",
	}
	fmt.Printf("%#v\n", a5)

	// nested structs initialization
	// p := Person{"John", "Doe", "", 45} // var v Person
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       45,
	}
	fmt.Printf("%#v\n", p)
	sam := Person{
		"Sam",
		"Smith",
		"000-10-0021",
		53,
		Address{street1: "2 2nd St.", street2: "Suite 101"},
	}
	fmt.Printf("%#v\n", sam)

	// anonymous fields
	c := City{"Brooklyn", State{"New York"}}
	fmt.Printf("%#v\n", c)
	fmt.Printf("City: %v, State: %v\n", c.string, c.name)
	fmt.Printf("City: %v, State: %v\n", c.string, c.State.name)

	// ambiguous fields in anonymous/embedded structs
	c1 := City2{
		Country: Country{"U.S.A"},
		string:  "Brooklyn", State: State{"New York"}}
	fmt.Printf("%#v\n", c1)
	fmt.Printf("City: %v, State: %v, Country: %v\n",
		c1.string, c1.State.name, c1.Country.name)
}
