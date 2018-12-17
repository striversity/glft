package input

import "testing"

type args struct {
	low int
	max int
}

func TestGetRandInt(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"invalid range 1", args{1, 1}, 0},
		{"invalid range 2", args{2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandInt(tt.args.low, tt.args.max); got != tt.want {
				t.Errorf("GetRandInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetRandInt2(t *testing.T) {
	tests2 := []struct {
		name string
		args args
	}{
		{"range 0, 0", args{0, 0}},
		{"range 0, 1", args{0, 1}},
		{"range 0, 2", args{0, 2}},
		{"range -1, 0", args{-1, 0}},
		{"range -1, 1", args{-1, 1}},
		{"range -19, 1", args{-19, 1}},
		{"range 10, 13", args{10, 13}},
		{"range 10, 13", args{10, 13}},
		{"range 10, 13", args{10, 13}},
		{"range 10, 13", args{10, 13}},
		{"range 10, 13", args{10, 13}},
		{"range 10, 13", args{10, 13}},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandInt(tt.args.low, tt.args.max); !inRange(tt.args, got) {
				t.Errorf("GetRandInt(%v, %v) = %v", tt.args.low, tt.args.max, got)
			}
		})
	}
}

func inRange(t args, n int) bool {
	if n >= t.low && n <= t.max {
		return true
	}

	return false
}
