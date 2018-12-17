package rw

import "io"
const(
	headerFmt = "[REC_HEADER_%v]"
	footerFmt = "[REC_FOOTER_%v]"
)
type (
	// RecordWriter wraps it input with a record 'header' and 'footer'.
	// The data between the 'header' and 'footer' is unchanged.
	RecordWriter struct {
		count  uint
		writer io.Writer // destination for records
	}
)