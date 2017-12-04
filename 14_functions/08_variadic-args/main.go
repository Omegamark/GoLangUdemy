package main

import "fmt"

func main() {
	// []float64 is != to float64
	data := []float64{43, 56, 87, 12, 45, 57}
	// ... means that the varialbe is a list and to take each item in the list.
	n := average(data...)
	fmt.Println(n)

}

// ...float64 means that this function can take any number of float64 arguments.
func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
