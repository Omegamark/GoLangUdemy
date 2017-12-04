package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println("I'm sf yo!", sf)
	fmt.Printf("%T \n", sf)
	// Javascript way of instantiating an accumulator
	// total := 0.0
	// Go way of instantiating an accumulator
	var total float64
	for _, v := range sf {
		// _ is a blank identifier for when a value for an operation is not needed.
		// an i variable in this location would contain the index.
		total += v
	}
	return total / float64(len(sf))
}
