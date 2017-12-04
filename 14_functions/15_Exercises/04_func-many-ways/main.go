package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 12, 3, 5, 10}
	funkManyWays(aSlice...)
}

func funkManyWays(z ...int) int {
	// this is the same solution as variadic function. I happened to solve exercise 5 before it was asked.
	var largestNum int
	for _, v := range z {
		if largestNum < v {
			largestNum = v
		}
	}
	fmt.Println(largestNum)
	return largestNum
}
