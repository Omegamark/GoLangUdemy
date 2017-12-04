package main

import "fmt"

func main() {

	name := "Mark"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name) // Donatello
}

func changeMe(z *string) {
	fmt.Println(z)  // The value of z here is the memory address of name since &name was passed to the function.
	fmt.Println(*z) // Mark
	*z = "Donatello"
	fmt.Println(z)  // The value at &name will change here since *z is now a pointer to the value stored at that specific address.
	fmt.Println(*z) // Donatello
}
