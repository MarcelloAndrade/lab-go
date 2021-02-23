package main

import (
	"fmt"
)

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	result := SimpleArraySum(ar)
	fmt.Println(result)
}

// SimpleArraySum https://www.hackerrank.com/challenges/simple-array-sum/problem
func SimpleArraySum(ar []int32) int32 {
	var sum int32
	for _, n := range ar {
		sum += n
	}
	return sum
}
