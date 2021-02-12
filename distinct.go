package main

import (
	"fmt"
	"sort"
)

func Distinct(A []int) int {

	sort.Ints(A)
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
