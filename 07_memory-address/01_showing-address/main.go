package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	// Printf - prints the desired thing with formatting expressed in the preceding string. Here I am telling Go to print out the address in decimal form and on a new line.
	fmt.Printf("%d \n", &a)
}
