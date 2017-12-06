package main

import (
	"fmt"
)

func main() {
	findDaVal()
}

func findDaVal() bool {
	var solution bool
	solution = (true && false) || (false && true) || !(false && false)
	fmt.Println(solution)
	return solution
}

// With (false && false) It returns false.
// With !(false && false) it returns true.
