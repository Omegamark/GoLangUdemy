package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 5)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println("Slice Length", len(mySlice))
	fmt.Println("Slice Capacity", cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
	fmt.Println(mySlice)
}
