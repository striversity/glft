package ms

import "testing"

func Test_Nil_MemStore_Write(t *testing.T) {
	var m *MemStore
	n, err := m.Write(nil)
	if n != 0 {
		t.Failed()
	}
	if err == nil {
		t.Failed()
	}
}

func Test_MemStore_Write(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		cap     uint
		args    args
		wantN   int
		wantErr bool
	}{
		{name: "0 bytes write", cap: 10, args: args{b: []byte{}}, wantN: 0, wantErr: false},
		{name: "nil b", cap: 10, args: args{b: nil}, wantN: 0, wantErr: false},
		{name: "5 bytes", cap: 5, args: args{b: []byte("hello")}, wantN: 5, wantErr: false},
		{name: "full buffer", cap: 4, args: args{b: []byte("hello")}, wantN: 4, wantErr: true},
	}
	for _, tt := range tests {
		m, _ := NewMemStore(tt.cap)
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := m.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemStore.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("MemStore.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
