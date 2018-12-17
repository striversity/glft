package db

import (
	"encoding/json"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetNextRecord(t *testing.T) {
	testName := "Get all records"
	t.Run(testName, func(t *testing.T) {
		numRecords, _:= loadTestData() // ignore load error, Load() test will handle that
		count := 0
		for {
			_, _, _, _, _, err := GetNextRecord()
			if err != nil {
				break
			}
			count++
		}

		if count != numRecords {
			t.Errorf("GetNextRecord() failed, got = %v, expected %v", numRecords, count)
		}
	})

}

func loadTestData() (num int, err error) {
	f, err := os.Open(jsonFile)
	if nil != err {
		return 0, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	d := []person{}
	err = dec.Decode(&d)
	if nil != err {
		return
	}
	_db.people = d
	_db.max = len(_db.people)
	log.Infof("%v Records loaded", _db.max)
	num = _db.max
	return
}
