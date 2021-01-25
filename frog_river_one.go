package main

import (
	"fmt"
)

func Solution7(X int, A []int) int {
	sum := (X * (X + 1)) / 2

	mapping := map[int]bool{}
	for i, v := range A {
		if _, ok := mapping[v]; !ok { // mapping 에 값이 없으면 초기화
			mapping[v] = true
			sum -= v
		}
		if sum == 0 {
			return i
		}
	}

	return -1
}

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(Solution7(5, A))
}
