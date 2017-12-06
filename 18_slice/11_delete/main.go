package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	// Appending from [2] all values from [3] on. So "Wedill not be printed
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
