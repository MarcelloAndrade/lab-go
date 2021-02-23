package main

import "fmt"

func main() {
	bigSum := VeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(bigSum)
}

// VeryBigSum https://www.hackerrank.com/challenges/a-very-big-sum/problem
func VeryBigSum(values []int64) int64 {
	var bigSum int64
	for _, value := range values {
		bigSum += value
	}
	return bigSum
}
