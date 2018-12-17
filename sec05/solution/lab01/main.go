// Section 05 - Exercise 01 : People Database Stats
package main

import (
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/sec05/solution/lab01/stats"
	"github.com/striversity/glft/sec05/solution/lab01/types"
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

	stats.Print(db)
}

func loadData(fn string) (db types.Db, err error) {
	var fr *input.FileReader
	if fr, err = input.NewFileReader(fn); err == nil {
		fr.ReadLine() // throw away header
		db = make(types.Db)
		var rec string
		for err == nil {
			rec, err = fr.ReadLine()
			if io.EOF == err { // deal with io.EOF differently
				return db, nil
			}
			if nil != err {
				break
			}
			p, _ := types.PersonFromCSV(strings.Trim(rec, "\n"))
			db[p.Gender] = append(db[p.Gender], p)
		}
	}
	return
}
func useage() {
	log.Errorf(`Usage: %v <file>
  where:
    file - is a CSV (comma-separated-value) formated text file`, os.Args[0])
}
