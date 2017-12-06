package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	// greeting[3] would be at at the 4th index of the slice, which exceeds the slice's length, causing an error.
	// The correct method is to append.
	greeting[3] = "Suprabadham"
	fmt.Println(greeting)

}
