package main

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Encryption 1", args{"haveaniceday"}, "hae and via ecy"},
		{"Encryption 2", args{"feedthedog"}, "fto ehg ee dd"},
		{"Encryption 3", args{"chillout"}, "clu hlt io"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.text); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
