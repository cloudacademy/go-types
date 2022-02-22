package main

import (
	"fmt"
)

func doubleNumber(number int) int {
	return 2 * number
}

func main() {
	fmt.Println()

	var y uint64 = 12345
	//var y uint64 = 1<<64 - 1

	var x int = doubleNumber(int(y))
	fmt.Println(x)
	fmt.Println()
}
