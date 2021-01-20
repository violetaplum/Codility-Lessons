package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(N int) int {
	count := 0
	binary := strconv.FormatInt(int64(N), 2)
	if strings.Count(binary, "1") == 1 {
		return 0
	}
	binaryArray := strings.Split(binary, "")
	zeroCount := strings.Split(binary, "1")
	countLan := len(zeroCount)
	lastIndex := len(binaryArray) - 1
	if binaryArray[lastIndex] == "0" {
		zeroCount = zeroCount[:countLan-1]
	}
	for _, v := range zeroCount {
		if count < len(v) {
			count = len(v)
		}
	}
	return count
}

func main() {
	fmt.Println(Solution(16))
}
