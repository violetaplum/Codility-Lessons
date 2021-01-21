package main

import "fmt"

func Solution4(X int, Y int, D int) int {
	locate := Y - X
	if locate%D != 0 {
		return locate/D + 1
	} else {
		return locate / D
	}
}

func main() {
	fmt.Println(Solution4(10, 85, 30))
}
