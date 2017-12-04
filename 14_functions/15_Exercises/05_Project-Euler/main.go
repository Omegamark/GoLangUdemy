package main

import (
	"fmt"
)

func main() {
	multiplesOfThreeAndFive(1000)
}

func multiplesOfThreeAndFive(z int) int {
	var solution int
	i := 0
	for i < z {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			solution += i
		}
		i++
	}
	fmt.Println(solution)
	return solution
}
