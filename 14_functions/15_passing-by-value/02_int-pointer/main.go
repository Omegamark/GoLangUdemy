package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // Stored at an 0x memory address

	changeMe(&age)

	fmt.Println(&age) // Same 0x memory address as above
	fmt.Println(age)  // 24
}

// The memory address (&age) is being passed by value to the changeMe() function. The changeMe() function receives it as a pointer * to and int (*int)
func changeMe(z *int) {
	fmt.Println(z)  // Same 0x memory address again
	fmt.Println(*z) // 44

	*z = 24

	fmt.Println(z)  // Same 0x memory addres yet again
	fmt.Println(*z) // 24
}

// Both age and z are variables pointing to the same memory address.
