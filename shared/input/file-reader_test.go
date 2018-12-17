package input

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name    string
		args    args
		wantFr  *FileReader
		wantErr bool
	}{
		{"non-existing file", args{"no-file.txt"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFr, err := NewFileReader(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFr, tt.wantFr) {
				t.Errorf("New() = %v, want %v", gotFr, tt.wantFr)
			}
		})
	}
}

func TestFileReader_ReadString(t *testing.T) {
	type fields struct {
		f *os.File
		s *bufio.Scanner
	}
	tests := []struct {
		name    string
		fields  fields
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FileReader{
				f: tt.fields.f,
				s: tt.fields.s,
			}
			gotS, err := fr.ReadLine()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileReader.ReadString(): error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("FileReader.ReadString(): got = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
