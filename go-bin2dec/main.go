package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var input string

func main() {
	fmt.Println("Input Binário:")
	fmt.Scanln(&input)

	if isValid(input) {
		fmt.Printf("Binário: %v Decimal: %v", input, binaryToDecimal(input))
	} else {
		fmt.Printf("Input inválido.")
	}
}

func isValid(input string) bool {
	matched, _ := regexp.MatchString("^[0-1]+$", input)
	if matched && (len(input) <= 8) {
		return true
	}
	return false
}

func binaryToDecimal(input string) float64 {
	decimal := 0.0
	s := strings.Split(input, "")
	for i := 0; i < len(s); i++ {
		base, _ := strconv.ParseFloat(s[i], 64)
		exponential := float64(len(s) - i - 1)
		p := math.Pow(2, exponential)
		decimal += base * p
		//fmt.Println("Base", base, "Binary: ", 2, "Exponential: ", exponential, "Pow: ", p, "Result: ", base*p)
	}
	return decimal
}
