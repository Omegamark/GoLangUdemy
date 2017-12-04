package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 12, 3, 5, 10}
	greatestNum(aSlice...)
}

func greatestNum(z ...int) int {
	var largestNum int
	for _, v := range z {
		if largestNum < v {
			largestNum = v
		}
	}
	fmt.Println(largestNum)
	return largestNum
}
