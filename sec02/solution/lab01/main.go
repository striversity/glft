/*
# Section 2 - Lab 1

Write a golang program which uses the provided package 'db'.
Find documenation for the db pacakge `go doc github.com/striversity/glft/shared/sec02/db`
and `go doc github.com/striversity/glft/shared/sec02/db GetNextRecord`. The 'db' package
provides functions to get records about people in a fictional database.

Use this information to answer the following:

## TODOS

1. TODO 1 - Who has the highest income?
2. TODO 2 - Who has the lowest income?
3. TODO 3 - What is the average income in the data set?
4. TODO 4 - Answer TODO 1, 2, 3 for women?
5. TODO 5 - Answer TODO 1, 2, 3 for men?

## REQUIREMENTS

1. Must use functions for demonstate common and re-usage code.
2. Must create at least two packages, including 'main'.
*/
package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/sec02/solution/lab01/stats"
	"github.com/striversity/glft/shared/sec02/db"
)

func main() {
  err:= db.Load() // must load data first
	if err != nil{
    log.Fatal(err)
  }
	fmt.Println("Stats for all records:")
	n, g, income := stats.HighestEarner(personRecord)
	fmt.Printf("Higest earner is %v (%v), with %v.\n", n, g, income)
	n, g, income = stats.LowestEarner(personRecord)
	fmt.Printf("Lowest earner is %v (%v), with %v.\n", n, g, income)
	income = stats.AvgIncome(personRecord)
	fmt.Printf("Average income is %v.\n", income)

	fmt.Println()
	fmt.Println("Stats for females:")
	n, _, income = stats.HighestEarner(filterGender("Female"))
	fmt.Printf("The higest earner amoung women is %v with %v.\n", n, income)
	n, _, income = stats.LowestEarner(filterGender("Female"))
	fmt.Printf("Lowest earner is %v with %v.\n", n, income)
	income = stats.AvgIncome(filterGender("Female"))
	fmt.Printf("Average income is %v.\n", income)

	fmt.Println()
	fmt.Println("Stats for Males:")
	n, _, income = stats.HighestEarner(filterGender("Male"))
	fmt.Printf("The higest earner amoung men is %v with %v.\n", n, income)
	n, _, income = stats.LowestEarner(filterGender("Male"))
	fmt.Printf("Lowest earner is %v with %v.\n", n, income)
	income = stats.AvgIncome(filterGender("Male"))
	fmt.Printf("Average income is %v.\n", income)
}


// return an anonymous function which closes over the parameter 'g'
func filterGender(g string) stats.RecordProvider {
	return func() (name, gender string, income db.Currency, err error) {
		var _fn, _ln, _g string
		var _income db.Currency

		for {
			_, _fn, _ln, _g, _income, err = db.GetNextRecord()
			if nil != err { // are we are the end?
				break
			}
			if g == _g {
				name = _ln + ", " + _fn
				gender = _g
				income = _income
				break
			}
		}
		return
	}
}

// adapt db.GetNextRecord() by wrapping it with another function
func personRecord() (name, gender string, income db.Currency, err error) {
	_, _fn, _ln, _g, _income, err := db.GetNextRecord()
	if nil != err { // are we are the end?
		return
	}

	name = _ln + ", " + _fn
	gender = _g
	income = _income
	return
}
