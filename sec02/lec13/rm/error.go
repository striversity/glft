package main

import (
	"fmt"
)

type error string

func (e error)Sting() string{
	return fmt.Sprintf("My Err: %v", string(e))
}
// New creates a custom error type
func New(s string) error{
	return error(s)
}
const ErrInf = error("Insufficient parameters")
const Err2 = error("Basic Err2")
const NoErr = error("")