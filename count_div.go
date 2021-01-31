package main

import "fmt"

func CountDiv(A int, B int, K int) int {
	// 세 정수 A,B,K 를 받는다
	// A~B 사이의 정수중 K 로 나뉘어지는 수를 리턴한다
	// A 는 항상 B 보다 작거나 같다
	// 예를 들어 A = 6 이고 B = 11 , K = 2  인 상황이라면 3을 반환해야한다

	Adiv := A / K
	Bdiv := B / K
	if A%K == 0 {
		return (Bdiv - Adiv) + 1
	} else {
		return Bdiv - Adiv
	}

}

func main() {
	A := 6
	B := 11
	K := 2
	fmt.Println(CountDiv(A, B, K))
}
