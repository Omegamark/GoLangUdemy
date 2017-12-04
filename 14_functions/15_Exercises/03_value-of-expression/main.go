package main

import (
	"fmt"
)

func main() {
	findDaVal()
}

func findDaVal() bool {
	var solution bool
	solution = (true && false) || (false && true) || (false && false)
	fmt.Println(solution)
	return solution
}

// It returns false.
