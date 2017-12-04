package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Mark", 36}
	fmt.Println(&c1.name) // 0x memory address

	changeMe(&c1)

	fmt.Println(c1)       // {Donatello 16}
	fmt.Println(&c1.name) // Same 0x memory address as above.
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Mark 36}
	fmt.Println(&z.name) // Same 0x memory address again
	z.name = "Donatello"
	z.age = 16
	fmt.Println(z)       // &{Donatello 16}
	fmt.Println(&z.name) // Same 0x memory address yet again
}
