package main

import "fmt"

func main() {
	funcManyWayz(1, 2)
	funcManyWayz(1, 2, 3)
	aSlice := []int{1, 2, 12, 3, 5, 10}
	funcManyWayz(aSlice...)
	funcManyWayz()
}

// Had to remove the expected return of int in order to run the program.

func funcManyWayz(z ...int) {
	fmt.Println(z)
}
