package main

import (
	"fmt"
)

func PassingCars(A []int) int {
	//TODO: A 배열을 받는다. A는 비지 않았다
	// A의 요소들은 연속적인 자동차들을 의미한다
	// A는 0이나 1을 가진다
	// - 0은 동쪽
	// - 1은 서쪽 을 향하는 자동차를 의미한다
	// (P,Q) 0 <= P < Q < N

	count := 0
	west := 0
	for _, v := range A {
		west += v
	}
	for _, v := range A {
		if v == 0 {
			count += west
			if 1000000000 < count {
				return -1
			}
		} else {
			west--
		}
	}

	return count
}

func main() {
	A := []int{0, 1, 0, 1, 1}
	fmt.Println(PassingCars(A))
}
