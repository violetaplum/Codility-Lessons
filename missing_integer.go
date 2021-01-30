package main

import (
	"fmt"
	"sort"
)

func MissingInteger(A []int) int {
	sort.Ints(A)
	if A[len(A)-1] <= 0 {
		// 마지막 인덱스가 0 이하면
		return 1
	}
	value := 0
	for _, v := range A {
		if v <= value {
			continue
		} else { // 커졌을때
			if v == value+1 {
				value++
				continue
			} else {
				return value + 1
			}
		}
	}
	return A[len(A)-1] + 1
}

func Solution9_ben(A []int) int {
	sort.Ints(A)
	value := 1
	for _, v := range A {
		if v == value {
			value++
		}
	}
	return value
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(MissingInteger(A))
}
