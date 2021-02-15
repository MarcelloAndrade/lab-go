package main

import (
	"fmt"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, n := range ar {
		sum += n
	}
	return sum
}

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	result := simpleArraySum(ar)
	fmt.Println(result)
}
