package rw

import (
	"errors"
)

func (rec *RecordWriter) Write(b []byte) (n int, err error) {
	if rec == nil {
		return n, errors.New("Invalid parameter, nil receiver")
	}

	// TODO 3 - complete implementation of this function

	return rec.writer.Write(b)
}
