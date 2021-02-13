package main

import (
	"fmt"
	"math"
	"sort"
)

func MaxProductOfTree(A []int) int {
	//TODO: 비지 않은 배열 A를 받았다
	// 이는 N 정수로 이루어져있다
	// (P, Q, R) 쌍이 있다
	// 이는 A[P] * A[Q] * A[R] 을 뜻한다
	// 0 <= P < Q < R < N
	// 저 세쌍을 완성하되 가장 큰 곱을 구해라
	if len(A) == 3 {
		return A[0] * A[1] * A[2]
	}
	//TODO: 일단 정렬시킨다
	// 모두 양수일경우엔 A[N-1] * A[N-2] * A[N-3]
	// 모두 음수일경우 A[N-1] * A[N-2] * A[N-3]
	// 혼합이면 A[N-1] * A[0] * A[1]
	// A[N-1] 마지막 항은 항상 곱해진다
	// A[N-1] 이 음수, 양수 여부에따라 최소값, 최대값을 구한다
	sort.Ints(A)
	last := A[len(A)-1]
	minimal := A[0] * A[1]
	max := A[len(A)-3] * A[len(A)-2]
	if last < 0 { // 모두다 음수일경우
		return max * last
	} else { // 모두다 양수이거나 혼합일경우
		return int(math.Max(float64(minimal), float64(max)) * float64(last))
	}
}

func main() {
	A := []int{-5, 5, -5, 4}
	fmt.Println(MaxProductOfTree(A))
}
