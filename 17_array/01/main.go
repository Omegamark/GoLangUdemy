package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	x[42] = string(777)
	// 42 is the character ' k ', 777 is the character ' ' '
	fmt.Println(x[42])
	fmt.Println(len(x))
	fmt.Println(x)
}
