package main

import "fmt"

func MaxCounters(N int, A []int) []int {
	// increase(X) - 카운터 X 을 1만큼 증가한다
	// max counter - 모든 카운터는 어떤 카운터든 가장 큰 값으로 정해진다
	// A 배열 은 M 정수를 가진다. 이 배열은 다음과같은 작업을 한다
	// 만약 A[k] = X ,  1<=X<=N 이면 K 오퍼레이션은 increase(X)이다
	// 만약 A[k] = N + 1 이면 K 오퍼레이션은 max counter 이다

	returnArray := make([]int, N)
	maxCounter := 0
	stacked := 0

	for _, v := range A {
		if v == N+1 {
			stacked = maxCounter
		} else if v >= 1 && v <= N {
			if returnArray[v-1] < stacked {
				returnArray[v-1] = stacked + 1
			} else if returnArray[v-1] >= stacked {
				returnArray[v-1]++
			}
			if maxCounter < returnArray[v-1] {
				maxCounter = returnArray[v-1]
			}
		}
	}

	for l, k := range returnArray {
		if k < stacked {
			returnArray[l] = stacked
		}
	}

	return returnArray
}

func Solution8_ben(N int, A []int) []int {
	result := map[int]int{}
	max := 0
	beforemax := 0
	for _, i := range A {
		if 1 <= i && i <= N {
			value, ok := result[i-1]
			if ok {
				if beforemax > value {
					result[i-1] = beforemax + 1
				} else {
					result[i-1] = value + 1
				}
			} else {
				result[i-1] = beforemax + 1
			}
			if max < result[i-1] {
				max = result[i-1]
			}

		} else if i == N+1 {
			beforemax = max
		}
	}

	total := []int{}
	for i := 0; i < N; i++ {
		value, ok := result[i]
		if ok {
			if beforemax < value {
				total = append(total, value)
				continue
			}
		}
		total = append(total, beforemax)
	}

	return total
}

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	fmt.Print(MaxCounters(5, A))
}
