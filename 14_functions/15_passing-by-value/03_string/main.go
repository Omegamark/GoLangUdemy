package main

import "fmt"

func main() {

	name := "Mark"
	fmt.Println(name) // Mark

	changeMe(name)

	fmt.Println(name) // Mark
	fmt.Println(&name)
}

func changeMe(z string) {
	fmt.Println(z) // Mark
	z = "Donatello"
	fmt.Println(z)  // Donatello
	fmt.Println(&z) // Address for &z is different from &name, which is why the value hasn't changed.
}
