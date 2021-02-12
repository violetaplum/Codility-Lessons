package main

import (
	"fmt"
	"sort"
)

func Distinct(A []int) int {
	sort.Ints(A)
	//TODO: 사실 여기서 sort는 필수요소는 아니다
	// map이기 때문에 어차피 정렬을 하나 안하나 똑같다
	mapmap := make(map[int]int, len(A))

	for _, v := range A {
		mapmap[v]++
	}
	return len(mapmap)
}

func main() {
	A := []int{2, 1, 1, 2, 3, 1}
	fmt.Println(Distinct(A))
}
