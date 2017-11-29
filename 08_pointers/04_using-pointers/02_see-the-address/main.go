package main

import "fmt"

// To clarify, imagine the variable in the zero function was z instead of x. These functions would still operate the same.
func zero(x int) {
	fmt.Printf("%p\n", &x) // address in func zero
	fmt.Println(&x)        // address in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
	// x is still 5 since Golang passes by value, so when x is passed
	// to the zero function, the variable x itself does not change. The
	// 2 x's are actually 2 different variables stored at 2 different memory
	// addresses. See the output.
}
