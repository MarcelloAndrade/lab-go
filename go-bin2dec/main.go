package main

import (
	"errors"
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

	decimal, err := binaryToDecimal(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Binário: ", input, " Decimal: ", decimal)
	}
}

func binaryToDecimal(input string) (float64, error) {
	// isValid ?
	matched, _ := regexp.MatchString("^[0-1]+$", input)
	if !matched || (len(input) > 8) {
		return 0, errors.New("Input inválido")
	}

	// convert
	var decimal float64
	s := strings.Split(input, "")
	for i := 0; i < len(s); i++ {
		base, _ := strconv.ParseFloat(s[i], 64)
		exponential := float64(len(s) - i - 1)
		p := math.Pow(2, exponential)
		decimal += base * p
		//fmt.Println("Base", base, "Binary: ", 2, "Exponential: ", exponential, "Pow: ", p, "Result: ", base*p)
	}
	return decimal, nil
}
