package main

import (
	"fmt"
	"strings"
)

func Rotation(A []int, K int) []int {
	if A == nil {
		return nil
	}
	if len(A) == 0 || len(A) == 1 || len(A) == K {
		return A
	}
	lenOfInt := len(A)
	if K%lenOfInt == 0 {
		return A
	}

	if len(A) == 2 {
		if K%2 != 1 {
			return A
		} else {
			return []int{A[1], A[0]}
		}
	}

	AStrList := strings.Split(strings.Trim(fmt.Sprint(A), "[]"), " ")
	lenIfStr := len(AStrList)
	AStr := strings.Trim(fmt.Sprint(A), "[]")
	count := strings.Count(AStr, AStrList[0])
	if count == lenIfStr {
		return A
	}

	returnArray := make([]int, lenOfInt)
	move := K % lenOfInt
	for i := 0; i < lenOfInt; i++ {
		index := (i + move) % lenOfInt
		returnArray[index] = A[i]
	}

	return returnArray
}

func Solution2_ben(A []int, K int) []int {
	if len(A) == 0 {
		return []int{}
	}
	result := A
	for i := 0; i < K; i++ {
		result = append([]int{result[len(result)-1]}, result[:len(result)-1]...)
	}
	return result
}
func main() {
	A := []int{}
	Rotation(A, 4)
}
