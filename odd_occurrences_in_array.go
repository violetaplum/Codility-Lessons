package main

import "sort"

func OddOccurrencesInArray(A []int) int {
	sort.Ints(A)
	var returnValue int
	if len(A) == 1 {
		returnValue = A[0]
	}
	i := len(A) - 1
	for {
		if i-1 <= 0 {
			break
		}
		if A[i] == A[i-1] {
			i -= 2
			continue
		} else {
			returnValue = A[i]
			break
		}
	}
	return returnValue
}

func Solution3_ben(A []int) int {
	data := map[int]int{}

	if len(A) == 1 {
		return A[0]
	}
	for _, i := range A {
		data[i] = data[i] + 1
	}

	for k, v := range data {
		if v%2 == 1 {
			return k
		}
	}

	return 0
}

func main() {
	A := []int{}
	OddOccurrencesInArray(A)
}
