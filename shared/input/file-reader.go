package input

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// FileReader maintains the state of a file opened for reading its data
// NOTE: FileReader will close the file automatically when it reaches the end
type FileReader struct {
	f *os.File
	s *bufio.Scanner
}

// NewFileReader creates a new FileReader for the file named by fn
func NewFileReader(fn string) (fr *FileReader, err error) {
	f, err := os.Open(fn)
	if nil != err {
		return
	}

	s := bufio.NewScanner(f)
	return &FileReader{f, s}, nil
}

// ReadLine returns a string from the file associated with FileReader
// until an error or end of file. io.EOF is returned AFTER the last line is read.
func (fr *FileReader) ReadLine() (s string, err error) {
	if nil == fr || nil == fr.s {
		return s, errors.New("Invalid receiver to FileReader.Read")
	}

	if !fr.s.Scan() {
		if nil == fr.s.Err() { // according to docs, this is EOF
			err = io.EOF
			fr.f.Close()
		}
		fr.f = nil
		fr.s = nil
		return
	}

	return fr.s.Text(), nil
}

// Scan returns true if it was able to read a line of text, false otherwise
func (fr *FileReader) Scan() bool {
	if nil == fr || nil == fr.s {
		return false
	}
	b := fr.s.Scan()
	if !b {
		if nil == fr.s.Err() { // according to docs, this is EOF
			fr.f.Close()
		}
		fr.f = nil
		fr.s = nil
	}
	return b
}

// Text returns the line of text read by Scan(). If end of file, then return ""
func (fr *FileReader) Text() string {
	if nil == fr || nil == fr.s {
		return ""
	}
	return fr.s.Text()
}
