package main

import "fmt"

func MinAvgTwoSlice(A []int) int {
	// A 배열에 N 정수들이 있다
	// (P, Q) 둘은 짝을 이루고 정수이다
	// 0 <= P < Q < N
	// A[P] + A[P+1} + ... + A[Q] 를 len(A)만큼 나누면 평균이 나온다
	// A[P]~A[Q] 까지의합 / (Q - P + 1)  이 평균이 된다는말이다
	// 짝의 갯수 -> 길이 까지 모든 수의 합 만큼 짝의 갯수가 정해진다

	//TODO: 2가지 케이스만 고려하면된다
	// slice가 2개 또는 3개인 경우이다. 평군의 성질로 부분집합의 평균은 가장 작은 인자보다 항상 크다.
	// 즉, (1,2) 의 평균은 1.5이므로 1보다 크다는 의미이다.
	// 또한 평균들의 평균은 각 인자들의 평균과 같다. 즉, (1,2,3,4)가 있을때, (1,2,3,4) = 2.5 이고
	// (1,2,) = 1.5 , (3,4) = 3.5 일때 (1.5, 3.5) = 2.5 가 된다는 의미
	// 위 두가지 성질을 생각했을때 인자의 수가 4개 이상인 것은 고려할 필요가 없다 가장 작은 평균을 가지는 부분집합은 가장 작은 숫자를 포함한
	// 2개 또는3개의 인자만 생각하면 된다
	// 3개의 인자를 고려하는 것은 2개의 부분집합으로는 3개의 부분집합을 구할 수 없기 때문이다. 가력 (1, 5, 2) 가 있을때,
	// (2,6,1) = 3 이고, (2,6) = 4, (6,1) = 3.5 일때 (4, 3.5) = 3.75 가 됨으로 3개의 경우는 따로 생각해야한다

	N := len(A)
	minAvg := float64(A[0]+A[1]) / 2
	minAvgStart := 0

	if N > 2 {
		first3Avg := float64(A[0]+A[1]+A[2]) / 3
		if first3Avg < minAvg {
			minAvg = first3Avg
		}
	}
	for i := 1; i < N-2; i++ {
		avg := float64(A[i]+A[i+1]) / 2
		if avg < minAvg {
			minAvg = avg
			minAvgStart = i
		}
		avg = float64(A[i]+A[i+1]+A[i+2]) / 3
		if avg < minAvg {
			minAvg = avg
			minAvgStart = i
		}
	}

	if N > 2 {
		last2Avg := float64(A[N-2]+A[N-1]) / 2
		if last2Avg < minAvg {
			minAvg = last2Avg
			minAvgStart = N - 2
		}
	}
	return minAvgStart
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(MinAvgTwoSlice(A))
}
