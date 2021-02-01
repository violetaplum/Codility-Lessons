package main

import (
	"fmt"
)

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	// DNA 시퀀스는 A,C,G,T 로 나타날 수 있다 이는 nucleotides 의 타입이 된다
	// 각각은 영향지수가 있고 이는 정수이다
	// A,C,G,T 는 각각 1,2,3,4 의 지수를 가진다
	// 주어진 DNA 시퀀스 중에 가장 작은 지수를 찾아라
	// P,Q 에는 M개의 쿼리가 있디
	// K 번째 쿼리에서 위치 P[K] 와 Q[K] 사이의 최소 지수를 찾아라
	M := len(P)
	sLen := len(S)
	aCnt := make([]int, sLen+1)
	cCnt := make([]int, sLen+1)
	gCnt := make([]int, sLen+1)
	min := make([]int, M)

	for i := 0; i < sLen; i++ {
		aCnt[i+1] = aCnt[i]
		cCnt[i+1] = cCnt[i]
		gCnt[i+1] = gCnt[i]
		if S[i] == 'A' {
			aCnt[i+1]++
		}
		if S[i] == 'C' {
			cCnt[i+1]++
		}
		if S[i] == 'G' {
			gCnt[i+1]++
		}
	}

	for j := 0; j < M; j++ {
		if aCnt[P[j]] != aCnt[Q[j]+1] {
			min[j] = 1
		} else if cCnt[P[j]] != cCnt[Q[j]+1] {
			min[j] = 2
		} else if gCnt[P[j]] != gCnt[Q[j]+1] {
			min[j] = 3
		} else {
			min[j] = 4
		}
	}

	return min
}

func main() {
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	S := "CAGCCTA"
	fmt.Println(GenomicRangeQuery(S, P, Q))
}
