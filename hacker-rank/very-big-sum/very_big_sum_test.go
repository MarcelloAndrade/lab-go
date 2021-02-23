package main

import "testing"

func TestVeryBigSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int64
		want   int64
	}{
		{"VeryBigSum 1", []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}, 5000000015},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VeryBigSum(tt.values); got != tt.want {
				t.Errorf("VeryBigSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
