package rw

import (
	"bytes"
	"io"
	"testing"
)

func TestNewRecordWriter(t *testing.T) {
	tests := []struct {
		name    string
		writer  io.Writer
		wantErr bool
	}{
		{name: "invalid parameter", writer: nil, wantErr: true},
		{name: "valid parameter", writer: &bytes.Buffer{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRecordWriter(tt.writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRecordWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
