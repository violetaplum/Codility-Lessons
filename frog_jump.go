package main

import "fmt"

func GrogJump(X int, Y int, D int) int {
	locate := Y - X
	if locate%D != 0 {
		return locate/D + 1
	} else {
		return locate / D
	}
}

func main() {
	fmt.Println(GrogJump(10, 85, 30))
}
