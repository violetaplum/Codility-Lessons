package main

import "fmt"

func Solution8(N int, A []int) []int {
	// increase(X) - 카운터 X 을 1만큼 증가한다
	// max counter - 모든 카운터는 어떤 카운터든 가장 큰 값으로 정해진다
	// A 배열 은 M 정수를 가진다. 이 배열은 다음과같은 작업을 한다
	// 만약 A[k] = X ,  1<=X<=N 이면 K 오퍼레이션은 increase(X)이다
	// 만약 A[k] = N + 1 이면 K 오퍼레이션은 max counter 이다
	returnArray := make([]int, N)
	maxCounter := 0

	for _, v := range A {
		if v == N+1 {
			for l, _ := range returnArray {
				returnArray[l] = maxCounter
			}
		} else if v >= 1 && v <= N {
			returnArray[v-1]++
			if maxCounter < returnArray[v-1] {
				maxCounter = returnArray[v-1]
			}
		}
	}

	return returnArray
}

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	fmt.Print(Solution8(5, A))
}
