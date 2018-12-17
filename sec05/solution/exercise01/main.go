// Section 05 - Exercise 01 : Countries with largest Cities
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/shared/input"
)

/*

 */

func main() {

	if len(os.Args) != 2 {
		useage()
		return
	}
	// load data from file into in-memory db
	db, err := loadData(os.Args[1])
	if err != nil {
		log.Errorf("Unable to read records: %v", err)
		useage()
		return
	}

	g := comparator(5)
	fmt.Println("Countries with 5 or more of the largest cities in the world:")
	for _, cities := range db {
		if g(cities) {
			fmt.Printf("\t%v\n", cities[0].Country)
		}
	}
}

// create a function that returns true if the number of cities is >= to x
func comparator(x int) func([]City) bool {
	return func(cities []City) bool {
		return len(cities) >= x
	}
}
func loadData(fn string) (db Db, err error) {
	var fr *input.FileReader
	if fr, err = input.NewFileReader(fn); err == nil {
		fr.ReadLine() // throw away header
		db = make(Db)
		var rec string
		for err == nil {
			rec, err = fr.ReadLine()
			if io.EOF == err { // deal with io.EOF differently
				return db, nil
			}
			if nil != err {
				break
			}
			c, _ := CityFromCSV(strings.Trim(rec, "\n"))
			db[c.Country] = append(db[c.Country], c)
		}
	}
	return
}
func useage() {
	log.Errorf(`Usage: %v <file>
  where:
    file - is a CSV (comma-separated-value) formated text file`, os.Args[0])
}
