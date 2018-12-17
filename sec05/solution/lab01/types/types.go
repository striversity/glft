package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	// Person is a struct with appropriate fields to represent an individual
	Person struct {
		FirstName string
		LastName  string
		SSN       string
		Gender    rune
		Age       uint8
		Salary    Currency
	}
	// Currency allows for the pretty printing of currency values
	Currency float64
	// People is a slice of Person
	People []Person
	// Db is a map of rune to People
	Db map[rune]People // map of gender to a people
)

// PersonFromCSV creates a Person object from a CSV string.
// Returns an error if 's' doesn't contain the exected fields
func PersonFromCSV(s string) (p Person, err error) {
	// sample: "Belia,Shrubsall,789-15-0853,F,65,526999.54"
	const expectedFields = 6
	fields := strings.Split(s, ",")
	if expectedFields > len(fields) {
		return p, errors.New("Incorrect number of fields in 's': " + s)
	}
	p.FirstName = fields[0]
	p.LastName = fields[1]
	p.SSN = fields[2]
	p.Gender = []rune(fields[3])[0] //cast to []rune and take first element
	age, _ := strconv.ParseUint(fields[4], 10, 8)
	p.Age = uint8(age)
	salary, _ := strconv.ParseFloat(fields[5], 64)
	p.Salary = Currency(salary)
	return p, nil
}
// String prints a Currency values using "$%1.2f", for example 5.49 as $5.49
func (c Currency) String() string {
	return fmt.Sprintf("$%1.2f", float64(c))
}
