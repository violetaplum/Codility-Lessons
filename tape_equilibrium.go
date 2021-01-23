package main

import (
	"fmt"
	"math"
)

func Solution6(A []int) int {
	length := len(A)
	if length < 1 {
		return 0
	}
	if length == 1 {
		if A[0] < 0 {
			return -A[0]
		} else {
			return A[0]
		}
	}
	if length == 2 {
		if (A[1] - A[0]) < 0 {
			return -(A[1] - A[0])
		} else {
			return A[1] - A[0]
		}
	}
	intint := make([][2]int, length)
	sum := 0
	lastSum := 0
	lastIndex := length - 1
	for i, v := range A {
		intint[i][0] = sum
		sum += v
		lastSum += A[lastIndex]
		intint[lastIndex][1] = lastSum
		lastIndex--
	}
	value := 0
	temp := 0
	if (intint[1][0] - intint[1][1]) < 0 {
		value = -(intint[1][0] - intint[1][1])
	} else {
		value = intint[1][0] - intint[1][1]
	}
	for i := 2; i < length; i++ {
		if (intint[i][0] - intint[i][1]) < 0 {
			temp = -(intint[i][0] - intint[i][1])
		} else {
			temp = intint[i][0] - intint[i][1]
		}
		if value > temp {
			value = temp
		}
	}
	return value
}

func Solution6_ben(A []int) int {
	result := -1
	if len(A) < 2 {
		return 0
	}
	total := 0
	left := 0
	for _, i := range A {
		total += i
	}
	for i := 0; i < len(A)-1; i++ {
		left += A[i]
		right := total - left
		val := left - right

		mindata := math.Abs(float64(val))
		if result == -1 {
			result = int(mindata)
		} else {
			result = int(math.Min(float64(result), mindata))
		}
	}

	return result
}

func main() {
	A := []int{-10, -20, -30, -40, 100}
	fmt.Println(Solution6(A))
}
