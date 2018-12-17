package db

import (
	"encoding/json"
  "errors"
  "fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/shared/input"
)

const (
	jsonFile   = "people.json"
	minRecords = 500
	maxRecords = 1000
)

type database struct {
	people []person
	cursor int
	max    int
}

var _db database

// Load Person records from people.json file into database, return nil on success or an error value
func Load() (err error) {
    // to ensure that we can always load the data, use the user's GOPATH

  dbFile := fmt.Sprintf(".%c%s", os.PathSeparator, jsonFile) 
	f, err := os.Open(dbFile)
	if nil != err {
		return errors.New("Can't open people.json. Did you copy people.json from shared/sec02/db to current directory?")
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	d := []person{}
	err = dec.Decode(&d)
	if nil != err {
		return err
	}
	_db.people = d[:input.GetRandInt(minRecords, maxRecords)]
	_db.max = len(_db.people)
	_db.cursor = 0
	log.Infof("%v Records loaded", _db.max)

	return
}
