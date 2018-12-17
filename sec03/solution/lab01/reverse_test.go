package main

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		s string
		want string
	}{
		{"Test 1", "Hello, world!", "!dlrow ,olleH"},
		{"Test 2", "Hello, 世界", "界世 ,olleH"},
		{"Test 3", "Todo está bien", "neib átse odoT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
