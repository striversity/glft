package db

// ResetCursor results in GetNextRecord() returrning the first person. Call ResetCursor() went GetNextRecord()
// returns an error marking no more records or to restart the iteration.
func ResetCursor() {
	_db.cursor = 0
}
