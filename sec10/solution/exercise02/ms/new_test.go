package ms

import (
	"testing"
)

func TestNewMemStore(t *testing.T) {
	type args struct {
		cap uint
	}
	tests := []struct {
		name      string
		args      args
		wantStore *MemStore
		wantErr   bool
	}{
		{"0 cap", args{0}, nil, false},
		{"10 cap", args{10}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStore, err := NewMemStore(tt.args.cap)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMemStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStore.cap != tt.args.cap {
				t.Failed()
			}
		})
	}
}
