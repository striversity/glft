package rw

import (
	"bytes"
	"errors"
	"fmt"
)

func (rec *RecordWriter) Write(b []byte) (n int, err error) {
	if rec == nil {
		return n, errors.New("Invalid parameter, nil receiver")
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(headerFmt, rec.count))
	buf.Write(b)
	buf.WriteString(fmt.Sprintf(footerFmt, rec.count))
	rec.count++
	
	return rec.writer.Write(buf.Bytes())
}
