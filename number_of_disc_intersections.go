package main

import (
	"fmt"
	"sort"
)

func NumberOfDiscIntersections(A []int) int {
	//TODO: 디스크는 각각 0~N 까지의 숫자를 가진다  (intersect : 교차하다)
	// 배열 A는 디스크들의 반지름을 가진다
	// J번째 디스크는 (J,0)번째 지점에 중심이 있게된다 그리고 A[J] 반지름을 가진다
	// J 디스크와 K 디스크가 J != K 라면 서로 교차한다
	// N은 0 ~ 100000 까지의 정수
	// A 내부 요소들은 0 ~ 2147483647 의 값을 가진다

	start := []int{}
	end := []int{}

	for i, v := range A {
		start = append(start, i-v)
		end = append(end, i+v)
	}

	sort.Ints(start)
	sort.Ints(end)

	intersection := 0
	index := 0
	for i := 0; i < len(A); i++ {
		for {
			if index >= len(A) || end[i] < start[index] {
				break
			}
			intersection += index
			intersection -= i
			index++
			fmt.Println("index ::", index)
			fmt.Println("i ::", i)
		}
	}

	if intersection > 10000000 {
		return -1
	}
	return intersection
}

func main() {
	A := []int{1, 5, 2, 1, 4, 0}
	fmt.Println(NumberOfDiscIntersections(A))
}
