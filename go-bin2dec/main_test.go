package main

import "testing"

func Test_binaryToDecimal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"Teste_Binary_0101101_to_Decimal_45", args{"0101101"}, 45, false},
		{"Teste_Binary_10101101_to_Decimal_173", args{"10101101"}, 173, false},
		{"Teste_Binary_10101_to_Decimal_21", args{"10101"}, 21, false},
		{"Teste_Error_Binary_10102_to_Decimal_0", args{"10102"}, 0, true},
		{"Teste_Error_Binary_010101010_to_Decimal_0", args{"010101010"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binaryToDecimal(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("binaryToDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binaryToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
