package main

import "testing"

func TestSimpleArraySum(t *testing.T) {
	tests := []struct {
		name string
		ar   []int32
		want int32
	}{
		{"SimpleArraySum 1", []int32{1, 2, 3, 4, 10, 11}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleArraySum(tt.ar); got != tt.want {
				t.Errorf("SimpleArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
