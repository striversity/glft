// Section 05 - Lecture 01 : Introduction to Structures
package main

import "fmt"

type (
	PersonHack []string // first name, last name, ssn, age
	FamilyHack []PersonHack
	DbHack     map[LastName]FamilyHack
	LastName   string
	Person     struct {
		firstName string
		lastName  string
		ssn       string
		age       uint8
	}
	Family []Person
	Db     map[LastName]Family
)

func main() {
	// without structs
	db := make(DbHack)
	mary := newPerson("Mary", "Smith", "000-00-0001", "21")
	john := newPerson("John", "Smith", "000-00-0002", "24")
	pete := newPerson("Pete", "Smith", "000-00-0003", "45")
	anne := newPerson("Anne", "Smith", "000-00-0004", "63")
	db["Smith"] = append(db["Smith"], mary, john, pete, anne)
	fmt.Printf("The Smith family: %v\n", db["Smith"])
	fmt.Printf("The Jones family: %v\n", db["Jones"])

	// life with structs
	var max Person
	max.firstName = "Max"
	max.lastName = "Jones"
	max.ssn = "001-00-0001"
	max.age = 30

	fmt.Printf("max: %v, %T\n", max, max)
	db2 := make(Db)
	mary2 := newPerson2("Mary", "Smith", "000-00-0001", 21)
	john2 := newPerson2("John", "Smith", "000-00-0002", 24)
	pete2 := newPerson2("Pete", "Smith", "000-00-0003", 45)
	anne2 := newPerson2("Anne", "Smith", "000-00-0004", 63)
	db2["Smith"] = append(db2["Smith"], mary2, john2, pete2, anne2)
	fmt.Printf("The Smith family: %v\n", db2["Smith"])
	db2["Jones"] = append(db2["Jones"], max)
	fmt.Printf("The Jones family: %v\n", db2["Jones"])
	// print entire person db
	fmt.Printf("Entire Person Db:\n%v\n", db2)
}
func newPerson2(fn, ln, ssn  string, age uint8) (p Person) {
	p.firstName =fn
	p.lastName =ln
	p.ssn =ssn
	p.age =age
	return
}
func newPerson(fn, ln, ssn, age string) (p PersonHack) {
	p = append(p, fn, ln, ssn, age)
	return
}
