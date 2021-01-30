package main

import (
	"fmt"
	"sort"
)

func PermCheck(A []int) int {
	// 배열 A가 순열인지 확인해라
	// A := []int{4,1,3} 이라면 2가 없으므로 순열이 아니다
	// 순열이면 1을 반환하고 아니면 0을 반환해라
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)
	for i, v := range A {
		if i+1 == v {
			continue
		} else {
			return 0
		}
	}

	return 1
}

func main() {
	A := []int{1}

	fmt.Println(PermCheck(A))
}
