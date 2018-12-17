package ms

import "testing"

func Test_Nil_MemStore_Read(t *testing.T) {
	var m *MemStore
	n, err := m.Read(nil)
	if n != 0 {
		t.Failed()
	}
	if err == nil {
		t.Failed()
	}
}

func Test_MemStore_Read(t *testing.T) {
	tests := []struct {
		name    string
		cap     uint
		in      []byte // data to write
		out     []byte // read into 'out'
		wantN   int
		wantErr bool
	}{
		{name: "0 bytes read", cap: 10, in: []byte("hello"), out: []byte{}, wantN: 0, wantErr: false},
		{name: "nil b read", cap: 10, in: []byte("hello"), out: nil, wantErr: false},
		{name: "empty buffer", cap: 10, in: nil, out: make([]byte, 10), wantN: 0, wantErr: true},
		{name: "5 bytes", cap: 10, in: []byte("hello"), out: make([]byte, 5), wantN: 5, wantErr: false},
		{name: "no more data", cap: 10, in: []byte("hello"), out: make([]byte, 6), wantN: 5, wantErr: true},
	}
	for _, tt := range tests {
		m, _ := NewMemStore(tt.cap)
		m.Write(tt.in)
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := m.Read(tt.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemStore.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("MemStore.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
