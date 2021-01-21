package main

import "fmt"

func Solution5(A []int) int {
	N := len(A) + 1
	sum := (N * (N + 1)) / 2
	sumOfA := 0
	for _, v := range A {
		sumOfA += v
	}
	return sum - sumOfA
}

func main() {
	fmt.Println(Solution5([]int{1}))
}
