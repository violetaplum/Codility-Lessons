package main

import (
	"fmt"
	"sort"
)

func Triangle(A []int) int {
	//TODO: A 는 정수 N 을 가진다
	// (P, Q, R) 이 있다면
	// 0 <= P < Q < R < N 관계를 가진다
	// 그리고
	// A[P] + A[Q] > A[R]
	// A[Q] + A[R] > A[P]
	// A[R] + A[P] > A[Q]
	// 는 삼각형의 조건이다
	// 삼각형이 하나라도 있으면 1, 없으면 0 을 리턴하

	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)
	max := A[len(A)-1]
	for i := len(A) - 2; i > 0; i-- {
		if A[i] > max/2 {
			if A[i] > (max - A[i-1]) {
				return 1
			} else {
				continue
			}
		} else {
			max = A[i]
			continue
		}
	}

	return 0
}

func main() {
	A := []int{10, 2, 5, 1, 8, 20}
	//B := []int{10,50,5,1}
	fmt.Println(Triangle(A))
	//fmt.Println(Triangle(B))
}
