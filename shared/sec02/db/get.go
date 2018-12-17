package db

import "errors"

/*
GetNextRecord returns several values corresponding to Peron's record.
Each call to this function returns a new record or non-nil error value
more records.
*** Important ***: Ensure the db is loaded first by calling Load(). */
func GetNextRecord() (id uint, firstName, lastName, gender string, income Currency, err error) {
	if _db.cursor < _db.max {
		r := _db.people[_db.cursor]
		// increment cursor
		_db.cursor++
		id = r.ID
		firstName = r.FirstName
		lastName = r.LastName
		gender = r.Gender
		income = r.Income
		return
	}
	err = errors.New("no more records, call ResetCuror() for a new iteration")
	return
}
