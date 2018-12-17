package wc

import (
	"testing"
)

func TestWriteCounter_String(t *testing.T) {
	tests := []struct {
		name string
		r    WriteCounter
		want string
		data []string
	}{
		{
			name: "0 writes",
			want: "There were 0 write operations totaling 0 bytes",
		},
		{
			name: "1 write",
			want: "There were 1 write operations totaling 0 bytes",
			data: []string{""},
		},
		{
			name: "5 write",
			want: "There were 5 write operations totaling 21 bytes",
			data: []string{"Hello", "world", "this", "is", "cool!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, b := range tt.data {
				tt.r.Write([]byte(b))
			}
			if got := tt.r.String(); got != tt.want {
				t.Errorf("WriteCounter.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteCounter_Write(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		r       *WriteCounter
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:    "nil test",
			wantErr: true,
		},
		{
			name:    "write test",
			r:       new(WriteCounter),
			args:    args{b: []byte("hello")},
			wantN:   5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.r.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteCounter.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
