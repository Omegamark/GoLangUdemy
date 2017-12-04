package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Mark":
		fmt.Println("Wassup Mark")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("You have no friends.")
	}
}
