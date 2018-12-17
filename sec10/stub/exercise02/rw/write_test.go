package rw

import (
	"bytes"
	"io"
	"testing"
)

func TestRecordWriter_Write(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name   string
		writer io.Writer
		args   []string
		want   string
	}{
		{name: "no data", writer: &bytes.Buffer{}, args: nil, want: "0 records written"},
		{name: "2 records", writer: &bytes.Buffer{}, args: []string{"Hi", "there"}, want: "2 records written"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &RecordWriter{
				writer: tt.writer,
			}
			for _, s := range tt.args {
				rec.Write([]byte(s))
			}

			got := rec.String()
			if got != tt.want {
				t.Errorf("RecordWriter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
