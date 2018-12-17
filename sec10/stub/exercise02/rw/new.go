package rw

import (
	"errors"
	"io"
)

// NewRecordWriter wraps an existing io.Writer such that records are delimited
func NewRecordWriter(w io.Writer) (rec io.Writer, err error) {
	// TODO 2 - complete implementation
}
