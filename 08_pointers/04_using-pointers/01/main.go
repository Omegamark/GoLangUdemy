package main

import "fmt"

// In these 2 functions, the memory address is being passed, referenced and manipulated,
// meaning that when the value of variable *z is changed at address &x.
func zero(z *int) {
	*z = 3
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
