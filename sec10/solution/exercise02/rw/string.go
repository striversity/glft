package rw

import "fmt"

func (rec RecordWriter) String() string {
	return fmt.Sprintf("%v records written", rec.count)
}
