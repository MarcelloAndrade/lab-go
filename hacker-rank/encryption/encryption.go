package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Printf(Encrypt("chillout"))
}

// Encrypt https://www.hackerrank.com/challenges/encryption/problem
func Encrypt(text string) string {
	var encrypted string

	textLen := len(text)
	textSqrt := math.Sqrt(float64(textLen))

	rows := int(math.Floor(textSqrt))
	columns := int(math.Ceil(textSqrt))

	if rows*columns < textLen {
		rows++
	}

	fmt.Println(rows, columns)
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			p := columns*r + c
			if p < textLen {
				encrypted += string(text[p])
			}
		}
		encrypted += " "
	}

	return strings.TrimSpace(encrypted)
}
